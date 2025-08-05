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
     role ENUM('userService','editor','admin') DEFAULT 'userService' COMMENT '角色',
     created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     updated_at  DATETIME  DEFAULT  CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
);
# 收藏表
CREATE TABLE IF NOT EXISTS favorites (
     id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
     user_id INT UNSIGNED COMMENT '',
     recipe_id INT UNSIGNED COMMENT '',
     created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     updated_at  DATETIME  DEFAULT  CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
);