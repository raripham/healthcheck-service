CREATE TABLE services (
	service_id SERIAL PRIMARY KEY,
    service_name VARCHAR(255),
    status VARCHAR(255),
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    uptime VARCHAR(255),
    service_metadata JSONB,
    node_id INT references nodes(node_id)
    );
   
CREATE TABLE nodes (
	node_id SERIAL PRIMARY KEY,
    node_name VARCHAR(255),
    node_ip VARCHAR(255),
    node_metadata JSONB
);


CREATE OR REPLACE FUNCTION update_end_time_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.end_time = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_end_time_column
BEFORE UPDATE ON services
FOR EACH ROW
EXECUTE FUNCTION update_end_time_column();


insert into services(service_name, status, start_time,end_time,uptime,node_id, service_metadata)
values ('noservice', 'Up', '2024-01-10 15:38:03.595 +0700', '2024-01-10 15:38:03.595 +0700', '00:00:00', '1', '{"env": "stg"}');

insert into nodes(node_name, node_ip, node_metadata)
values ('prod-ftech', '192.168.1.3', '{"env": "prod"}') returning node_id;