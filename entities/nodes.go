package entities

import (
	"database/sql"
	"healthcheck/model"
	"log"
)

func MapperNodeTest(data *sql.Rows) []model.ListNodeRequest {
	var nodes []model.ListNodeRequest
	for data.Next() {
		var no model.ListNodeRequest
		err := data.Scan(&no.NodeId, &no.NodeName, &no.NodeIp, &no.NodeMetadata)
		if err != nil {
			log.Fatal(err)
		}
		nodes = append(nodes, no)
	}
	return nodes
}

func NodeRegisterTest(db *sql.DB, no model.AddNodeRequest) int64 {
	var nodeId int64
	// check exist
	dataNo, err := db.Query(`SELECT * FROM nodes WHERE node_name = $1 and node_ip = $2;`, no.NodeName, no.NodeIp)
	if err != nil {
		log.Fatal(err)
	}
	nodes := MapperNodeTest(dataNo)
	if nodes != nil {
		nodeId = nodes[0].NodeId
	} else {
		noId, err := db.Query(`INSERT INTO nodes(node_name, node_ip, node_metadata) VALUES ($1, $2, $3) returning node_id`, no.NodeName, no.NodeIp, no.NodeMetadata)
		if err != nil {
			log.Fatal(err)
		}
		noId.Scan(&nodeId)
		noId.Close()
	}
	dataNo.Close()
	return nodeId
}
