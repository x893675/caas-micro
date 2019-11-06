INSERT INTO user.g_menu (id, created_at, updated_at, deleted_at, record_id, name, sequence, icon, router, hidden, parent_id, parent_path, creator) VALUES (1, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', '用户管理', 1170000, 'user', '/system/user', 0, '', '', '');


INSERT INTO user.g_menu_action (id, created_at, updated_at, deleted_at, menu_id, code, name) VALUES (1, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', 'add', '新增');
INSERT INTO user.g_menu_action (id, created_at, updated_at, deleted_at, menu_id, code, name) VALUES (2, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', 'edit', '编辑');
INSERT INTO user.g_menu_action (id, created_at, updated_at, deleted_at, menu_id, code, name) VALUES (3, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', 'del', '删除');
INSERT INTO user.g_menu_action (id, created_at, updated_at, deleted_at, menu_id, code, name) VALUES (4, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', 'query', '查询');


INSERT INTO user.g_menu_resource (id, created_at, updated_at, deleted_at, menu_id, code, name, method, path) VALUES (1, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', 'query', '查询用户数据', 'GET', '/api/v1/users');
INSERT INTO user.g_menu_resource (id, created_at, updated_at, deleted_at, menu_id, code, name, method, path) VALUES (2, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', 'get', '精确查询用户数据', 'GET', '/api/v1/users/:id');
INSERT INTO user.g_menu_resource (id, created_at, updated_at, deleted_at, menu_id, code, name, method, path) VALUES (3, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', 'create', '创建用户数据', 'POST', '/api/v1/users');
INSERT INTO user.g_menu_resource (id, created_at, updated_at, deleted_at, menu_id, code, name, method, path) VALUES (4, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', 'update', '更新用户数据', 'PUT', '/api/v1/users/:id');
INSERT INTO user.g_menu_resource (id, created_at, updated_at, deleted_at, menu_id, code, name, method, path) VALUES (5, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', 'delete', '删除用户数据', 'DELETE', '/api/v1/users/:id');
INSERT INTO user.g_menu_resource (id, created_at, updated_at, deleted_at, menu_id, code, name, method, path) VALUES (6, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', 'disable', '禁用用户数据', 'PATCH', '/api/v1/users/:id/disable');
INSERT INTO user.g_menu_resource (id, created_at, updated_at, deleted_at, menu_id, code, name, method, path) VALUES (7, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', 'enable', '启用用户数据', 'PATCH', '/api/v1/users/:id/enable');
INSERT INTO user.g_menu_resource (id, created_at, updated_at, deleted_at, menu_id, code, name, method, path) VALUES (8, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', 'queryRole', '查询角色数据', 'GET', '/api/v1/roles');

INSERT INTO user.g_role (id, created_at, updated_at, deleted_at, record_id, name, sequence, memo, creator) VALUES (1, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, '0b4628a5-683a-4b15-8c14-0448a0737945', 'userAdmin', 1170000, '用户管理员', '');
INSERT INTO user.g_role (id, created_at, updated_at, deleted_at, record_id, name, sequence, memo, creator) VALUES (2, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, '0b4628a5-683a-4b15-8c14-0448a0737946', 'userCommon', 1170000, '普通用户', '');

INSERT INTO user.g_role_menu (id, created_at, updated_at, deleted_at, role_id, menu_id, action, resource) VALUES (1, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, '0b4628a5-683a-4b15-8c14-0448a0737945', 'd5c6bf94-a1c2-4328-b7d8-c2c99fa639f2', 'add,edit,query,del', 'query,get,create,update,delete,enable,disable,queryRole');

INSERT INTO user.g_user (id, created_at, updated_at, deleted_at, record_id, user_name, real_name, password, email, phone, status, creator) VALUES (1, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, '58665369-8ae1-43ea-b540-6a102c134c94', 'admin', 'admin', 'd033e22ae348aeb5660fc2140aec35850c4da997', 'admin@admin.com', '', 1, '');
INSERT INTO user.g_user (id, created_at, updated_at, deleted_at, record_id, user_name, real_name, password, email, phone, status, creator) VALUES (2, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, '79f49f9d-79c7-4583-a80e-22ea4a8b44c3', 'developer', 'developer', '3dacbce532ccd48f27fa62e993067b3c35f094f7', 'developer@developer.com', '', 1, '');

INSERT INTO user.g_user_role (id, created_at, updated_at, deleted_at, user_id, role_id) VALUES (1, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, '58665369-8ae1-43ea-b540-6a102c134c94', '0b4628a5-683a-4b15-8c14-0448a0737945');
INSERT INTO user.g_user_role (id, created_at, updated_at, deleted_at, user_id, role_id) VALUES (2, '2019-11-04 15:53:29', '2019-11-04 15:53:29', null, '79f49f9d-79c7-4583-a80e-22ea4a8b44c3', '0b4628a5-683a-4b15-8c14-0448a0737946');

