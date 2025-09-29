# 创建 recipes 表
CREATE TABLE IF NOT EXISTS recipes (
       id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '菜谱ID',
       name VARCHAR(100) NOT NULL COMMENT '菜名',
       description TEXT COMMENT '简要描述',
       instructions  TEXT NOT NULL COMMENT '做法步骤（支持 Markdown 或 HTML）',
       prep_time  INT COMMENT '准备时间（分钟）',
       cook_time  INT COMMENT '烹饪时间（分钟）',
       servings  INT  COMMENT  '分量（几人份）',
       image_url VARCHAR(255)  COMMENT '封面图 URL',
       created_at  DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
       updated_at  DATETIME  DEFAULT  CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
       status  ENUM('draft','published','archived') NOT NULL DEFAULT 'published' COMMENT '状态'
);
# 代表大的菜系，如“川菜”、“湘菜”、“粤菜”等。
CREATE TABLE IF NOT EXISTS categories (
      id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '分类 ID',
      name VARCHAR(50) NOT NULL UNIQUE COMMENT '分类名',
      description TEXT COMMENT '描述',
      parent_id INT UNSIGNED COMMENT '父级分类 ID',
      level TINYINT DEFAULT 1 COMMENT '层级',
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
      updated_at  DATETIME  DEFAULT  CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
);
# 用于管理所有类型的标签，通过 type 字段区分不同类别
CREATE TABLE IF NOT EXISTS tags (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY	COMMENT '标签ID',
    name VARCHAR(50) NOT NULL UNIQUE COMMENT '标签名，如“辣”、“春节”、“意大利”',
    type ENUM('cuisine', 'country', 'flavor', 'holiday', 'ingredient', 'other') NOT NULL COMMENT '标签类型',
    description  TEXT	COMMENT '描述',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME  DEFAULT  CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
);
# 多对多关系：一个菜谱可以有多个标签，一个标签可用于多个菜谱
CREATE TABLE IF NOT EXISTS recipe_tags (
   id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY	COMMENT 'id',
   recipe_id  BIGINT UNSIGNED COMMENT '菜谱 ID',
   tag_id  INT UNSIGNED COMMENT '标签 ID'
);
# 食材主表
CREATE TABLE IF NOT EXISTS ingredients (
   id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '食材 ID',
   name VARCHAR(100) NOT NULL UNIQUE COMMENT '食材名',
   category VARCHAR(50)  COMMENT '分类',
   unit VARCHAR(20) COMMENT '常用单位',
   created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   updated_at  DATETIME  DEFAULT  CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
);
# 记录每道菜中使用的食材及用量
CREATE TABLE IF NOT EXISTS recipe_ingredients (
      id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
      recipe_id BIGINT UNSIGNED COMMENT '菜谱ID',
      ingredient_id INT UNSIGNED COMMENT '食材ID',
      quantity DECIMAL(10,2) COMMENT '用量',
      unit VARCHAR(20) COMMENT '单位（可覆盖食材默认单位）',
      is_main BOOLEAN DEFAULT FALSE COMMENT '是否为主要食材',
      sort_order INT DEFAULT 0 COMMENT '排序',
      INDEX idx_recipe_ingredient_id (recipe_id,ingredient_id)
);
# 如果支持用户投稿或收藏
CREATE TABLE IF NOT EXISTS users (
     id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '用户 ID',
     username VARCHAR(50) NOT NULL  UNIQUE  COMMENT '用户名',
     email VARCHAR(100) UNIQUE COMMENT '邮箱',
     password_hash VARCHAR(255) NOT NULL COMMENT '密码哈希',
     role int NOT NULL DEFAULT 2 COMMENT '权限',
     created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     updated_at  DATETIME  DEFAULT  CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
     deleted_at DATETIME DEFAULT NULL COMMENT '删除时间'
);
# 收藏表
CREATE TABLE IF NOT EXISTS favorites (
     id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
     user_id INT UNSIGNED COMMENT '',
     recipe_id INT UNSIGNED COMMENT '',
     created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     updated_at  DATETIME  DEFAULT  CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
);

# 流程表
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

# 步骤表
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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