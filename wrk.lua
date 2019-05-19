wrk.method = "POST"
wrk.body   = '[ { "eventId": "5cb8afb8-5c81-44be-ab9b-167764ff4e71", "data": { "userId": "0a41c6e5-5ad9-4412-ac5c-fc4cc13e2d6c", "cartId": "a40ea1f4-b2f0-4904-9a3c-85c27efeca92", "productId": "5ca94d88036709a43a429a2b", "numItems": 9, "pricePerItem": 99.23, "productOptions": {"color": "blue"} } }, { "eventId": "20027f1d-8d68-455b-99ea-39b5fa5495a0", "data": { "userId": "5c564d55-5393-48d4-94f5-0ed3402f2a08", "cartId": "ed3ffd05-11cd-41c8-ab58-69292f9a0a9f", "productId": "5ca94e34f7e55570ecb918af", "numItems": 2, "pricePerItem": 93.82, "productOptions": {"color": "red"} } } ]'

wrk.headers["Content-Type"] = "application/json"
wrk.headers["Accept"] = "application/json"