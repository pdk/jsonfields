# jsonfields

get the names of the fields in a JSON document

    $ echo '[{"a": 12, "b": "ack"}]' | jsonfields
    a b
    $ echo '[{"a": 12, "b": "ack"}]' | jsonfields -c
    a: 1
    b: 1
