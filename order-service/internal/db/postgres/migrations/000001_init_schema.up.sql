CREATE TABLE ordered_devices
(
    order_uuid  TEXT,
    device_uuid TEXT,
    amount      int4 CHECK (amount > 0)
);


CREATE TABLE orders
(
    order_uuid  TEXT,
    user_uuid   TEXT,
    order_price float CHECK ( order_price > 0),
    status      int,
    created_at  timestamptz
)