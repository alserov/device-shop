CREATE TABLE orders
(
    user_uuid   TEXT,
    order_uuid  TEXT,
    device_uuid TEXT,
    amount      int4 CHECK (amount > 0),
    status      int,
    created_at  timestamptz
)