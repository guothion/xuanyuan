# 设计一个 Process
我们想设计一个审核process 的系统。
表设计：

| 状态             | 	说明                   |
|----------------|-----------------------|
| **draft**      | 	草稿：流程已创建但未提交（用户正在编辑） |
| **pending**    | 	进行中：已提交，至少有一个步骤待处理   |
| **approved**	  | 已通过：所有步骤完成且全部通过       |
| **rejected**   | 	已拒绝：某个步骤被拒绝，流程终止     |
| **canceled**   | 	已取消：由发起人或管理员主动取消     |
| **expired**    | 已过期：长时间未处理自动关闭（可选）    |
| **terminated** | 	已终止：系统强制终止（如违规）      |
- ## process
```sql
CREATE TABLE process (
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    code         VARCHAR(50) UNIQUE NOT NULL COMMENT '流程唯一编码，如 leave_process',
    biz_id       VARCHAR(64) COMMENT '关联业务ID',  -- 关联业务ID（如报销单ID）
    name         VARCHAR(100) NOT NULL COMMENT '流程名称，如 请假审批流程',
    description  TEXT NULL COMMENT '流程描述',
    status       ENUM('draft', 'pending', 'approved', 'rejected', 'canceled', 'expired', 'terminated') DEFAULT 'draft' COMMENT '状态',
    created_at   DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_code (code),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```
- ## process_steps

| 状态             | 	说明                             |
|----------------|---------------------------------|
| **pending**    | 	待处理：尚未开始，等待处理人操作               |
| **processing** | 	处理中：已领取/开始处理，但未提交结果（可选）        |
| **approved**   | 	已通过：该步骤已批准                     |
| **rejected**   | 	已拒绝：该步骤被拒绝                     |
| **skipped**    | 	已跳过：因条件不满足或流程跳转被跳过（如金额小自动跳过财务） |
| **canceled**   | 	已取消：该步骤被取消（如流程撤回）              |

```sql
CREATE TABLE process_steps (
    id              BIGINT AUTO_INCREMENT PRIMARY KEY,
    process_code    VARCHAR(50) NOT NULL COMMENT '关联的流程 process_code',
    step_key        VARCHAR(50) NOT NULL COMMENT '步骤唯一标识，如 submit, approve_hr, approve_mgr',
    name            VARCHAR(100) NOT NULL COMMENT '步骤显示名称，如 提交申请、HR 审批',
    type            VARCHAR(30) NOT NULL COMMENT '步骤类型：approval, notification, auto_task, decision 等',
    status          ENUM('pending', 'processing', 'approved', 'rejected', 'skipped', 'canceled') DEFAULT 'pending' COMMENT '节点状态',
    sort_order      INT NOT NULL DEFAULT 0 COMMENT '排序序号，用于确定执行顺序',
    required        TINYINT(1) DEFAULT 1 COMMENT '是否必须完成',
    description     TEXT NULL COMMENT '描述',
    
    -- 扩展字段（可根据业务需要添加）
--     assignee_id    INT DEFAULT 'role' COMMENT '处理人id',
--     assignee_name  VARCHAR(50) COMMENT '处理人姓名',
    assignees      VARCHAR COMMENT '可以处理当前节点的人',
    duration_limit  INT COMMENT '处理时限（分钟），超时可提醒',
    
    -- 标签支持（可用于前端渲染、条件判断）
    tags            JSON COMMENT '标签数组，如 ["urgent", "finance", "skip_if_weekend"]',
    config          JSON COMMENT '额外配置，如表单字段、条件表达式等',
    
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY uk_process_step (process_code, step_key),
    INDEX idx_process_code (process_code),
    INDEX idx_type (type),
    INDEX idx_assignee (assignee_type, assignee_value)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

- ## process_logs
```sql
CREATE TABLE process_logs (
    id              BIGINT AUTO_INCREMENT PRIMARY KEY,
    
    -- 关联信息
    process_id     BIGINT NOT NULL COMMENT '流程实例 ID',
    step_key        VARCHAR(50) NOT NULL COMMENT '操作的步骤 key',
    step_name       VARCHAR(100) NOT NULL COMMENT '步骤显示名称，冗余便于查询',
    
    -- 操作人信息
    operator_id     BIGINT NOT NULL COMMENT '操作人用户 ID',
    operator_name   VARCHAR(50) NOT NULL COMMENT '操作人姓名，冗余存储避免关联',
    operator_role   VARCHAR(50) COMMENT '操作人角色，如 dept_manager, hr',
    
    -- 操作类型
    action          ENUM('submit', 'approve', 'reject', 'recall', 'cancel', 'skip', 'delegate', 'timeout') NOT NULL COMMENT '操作类型',
    
    -- 内容与描述
    comment         TEXT COMMENT '操作备注或审批意见',
    description     TEXT COMMENT '系统自动生成的描述，如“系统因超时自动拒绝”',
    
    -- 状态变更（可选，用于快速查询）
    from_status     VARCHAR(20) COMMENT '操作前状态，如 pending → approved',
    to_status       VARCHAR(20) COMMENT '操作后状态',
    
    -- 元数据
    client_ip       VARCHAR(45) COMMENT '操作 IP 地址',
    user_agent      TEXT COMMENT '客户端信息',
    extra_data      JSON COMMENT '额外数据，如委派人、跳转规则等',
    
    -- 时间
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
    
    -- 索引
    INDEX idx_process_id (process_id),
    INDEX idx_action (action),
    INDEX idx_operator (operator_id),
    INDEX idx_created_at (created_at),
    INDEX idx_step_key (step_key)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='流程操作日志表';
```