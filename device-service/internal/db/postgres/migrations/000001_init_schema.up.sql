CREATE TABLE devices
(
    UUID         TEXT,
    Title        TEXT,
    Description  TEXT,
    Price        float8 CHECK (Price > 0),
    Manufacturer TEXT,
    Amount       int CHECK (Amount > 0)
)