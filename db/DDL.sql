-- auto-generated definition
create table kitchen_order
(
    id          bigint unsigned auto_increment comment '自增id' primary key,
    unique_code varchar(128)  default ''                not null comment '订单号[全局唯一]',
    create_id   varchar(128)  default ''                not null comment '创建人id',
    create_name varchar(128)  default ''                not null comment '创建人name',
    info        varchar(2048) default ''                not null comment '订单内容，暂时不做索引直接怼进去',
    status      int           default 1                 not null comment '订单状态， 1-新建[提单]  2-已支付',
    target_time timestamp     default CURRENT_TIMESTAMP not null comment '订单的目标时间[精确到分钟]',
    extra       varchar(2048) default ''                not null comment '扩展字段',
    create_time timestamp     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time timestamp     default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '最后更新时间',
    constraint uk_order_code
        unique (unique_code),
    index idx_order_creator (create_id, target_time)
) comment '餐厅订单表[状态表]';

create table user_info
(
    id            bigint unsigned auto_increment comment '自增id' primary key,
    openId        varchar(128) default ''                not null comment '微信提供的openId',
    user_nickName varchar(128) default ''                not null comment '创建人name',
    user_icon     MEDIUMBLOB comment '暴力存储头像文件',
    status        int          default 1                 not null comment '1-正常状态  2-删号',
    create_time   timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time   timestamp    default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '最后更新时间',
    constraint uk_open_id
        unique (openId)
) comment '用户信息';
