package beatmap_meili

import (
	"encoding/json"
	"keeper/internal/entity"
	rep "keeper/internal/repository"

	"github.com/meilisearch/meilisearch-go"
)

type repository struct {
	meili *meilisearch.Client
}

func New(meili *meilisearch.Client) rep.BeatmapMeiliRepository {
	return &repository{meili: meili}
}

func (r *repository) UpdateBeatmap(u *entity.BeatmapSet) error {
	index := r.meili.Index("beatmaps_full")
	var data map[string]interface{}
	json_data, _ := json.Marshal(u)
	json.Unmarshal(json_data, &data)
	_, err := index.UpdateDocuments(data)
	return err
}

// func (r *repository) GetUserByID(id int) (*entity.User, error) {
// 	index := r.meili.Index("users")

// 	res, err := index.Search("", &meilisearch.SearchRequest{
// 		Filter: fmt.Sprintf("id = %d", id),
// 		Limit:  1,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	b, err := json.Marshal(res.Hits[0])
// 	if err != nil {
// 		return nil, err
// 	}
// 	result := entity.User{}
// 	err = json.Unmarshal(b, &result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }
