package redis

import (
	"fmt"
	"strings"

	"github.com/tonychill/ifitu/apis/pb/go/global"
	"github.com/xtgo/uuid"
)

func SetRedisKey(index, id string) string {
	return fmt.Sprintf("%s:%s", index, id)

}

func getRawQuery(query *global.Query) (q string, err error) {
	if query == nil {
		return "", fmt.Errorf("query is nil")

	}
	var (
		keyTracker = make(map[string]int)
	)

	sb := strings.Builder{}
	for _, t := range query.Terms {
		if _, ok := keyTracker[t.Key]; ok {
			// TODO: need to better handle cases where the user passes multiple user ids on queries.
			//not that there is a fix needed but this should be improved.
			continue
		}
		keyTracker[t.Key] = keyTracker[t.Key] + 1
		// Handle uuids for redis search
		cleanString := ""
		_, err := uuid.Parse(t.Value)
		if err == nil {
			cleanString = strings.ReplaceAll(t.Value, "-", "")
		} else {
			cleanString = t.Value
		}

		_, err = sb.WriteString("@" + t.Key + ":" + cleanString + " ")
		if err != nil {
			return "", fmt.Errorf(
				"failed to write property (%s) and value (%s) to buffer | errMsg: %s",
				t.Key, t.Value, err.Error())
		}

	}

	if query.Page == 0 {
		//
	} else {
		// add the offset to the query
		//  page * resultsPerPage
	}

	// TODO: add in pagination, sorting, and any other query params that may be included in the query
	// log.Debug().Msgf("redis query: %s", sb.String())

	return sb.String(), err
}
