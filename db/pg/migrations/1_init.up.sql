create table lists
(
    id    uuid primary key,
    title varchar(50)
);

create table items
(
    id          uuid primary key,
    list_id     uuid,
    description varchar(255)
);

create table items_lists
(
    item_id uuid,
    list_id uuid,
    primary key (item_id, list_id),
    foreign key (item_id) references items (id),
    foreign key (list_id) references lists (id)
);
