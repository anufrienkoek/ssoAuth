INSERT INTO apps (id, name, secret)
VALUES (1, 'TEST', 'test-secret')
ON CONFLICT DO NOTHING;