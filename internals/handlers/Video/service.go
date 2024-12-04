package video

import (
	"context"

	"github.com/salman-aziz-4425/Trello-reimagined/internals/db"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/dtos"
)

func AddGroupLogic(videoDetails dtos.VideoGroupDTO) error {
	println("AddGroup")
	println("User:", videoDetails.AuthorId)
	println("Video Group Name:", videoDetails.GroupName)
	_, err := db.Pool.Exec(
		context.Background(),
		"INSERT INTO videos (author_id, title) VALUES ($1, $2)",
		videoDetails.AuthorId,
		videoDetails.GroupName,
	)

	if err != nil {
		println("Error inserting into videos table:", err.Error())
		return err
	}
	println("Video group added successfully!")
	return nil
}

func GetGroupDetailsLogic(authorId int) ([]struct {
	ID        int
	AuthorId  int
	GroupName string
}, error) {
	println("getGroupDetails")
	rows, err := db.Pool.Query(
		context.Background(),
		"SELECT id as ID,title,author_id FROM videos WHERE author_id = $1",
		authorId,
	)
	if err != nil {
		println("Error querying videos table:", err.Error())
		return nil, err
	}
	defer rows.Close()

	var groups []struct {
		ID        int
		AuthorId  int
		GroupName string
	}
	for rows.Next() {
		var group struct {
			ID        int
			AuthorId  int
			GroupName string
		}
		err := rows.Scan(&group.ID, &group.GroupName, &group.AuthorId)
		if err != nil {
			println("Error scanning row:", err.Error())
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
