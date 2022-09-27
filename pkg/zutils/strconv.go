/* Copyright 2022 Zinc Labs Inc. and Contributors
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package zutils

import (
	"fmt"
	"strconv"
	"strings"
)

func ToString(v interface{}) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case int:
		return strconv.Itoa(v), nil
	case bool:
		return strconv.FormatBool(v), nil
	default:
		return fmt.Sprintf("%v", v), nil
	}
}

func ToFloat64(v interface{}) (float64, error) {
	switch v := v.(type) {
	case float64:
		return v, nil
	case uint64:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case int:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, fmt.Errorf("ToFloat64: unknown supported type %T", v)
	}
}

func ToUint64(v interface{}) (uint64, error) {
	switch v := v.(type) {
	case uint64:
		return v, nil
	case float64:
		return uint64(v), nil
	case int64:
		return uint64(v), nil
	case int:
		return uint64(v), nil
	case string:
		return strconv.ParseUint(v, 10, 64)
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, fmt.Errorf("ToInt: unknown supported type %T", v)
	}
}

func ToInt(v interface{}) (int, error) {
	switch v := v.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	case uint64:
		return int(v), nil
	case int64:
		return int(v), nil
	case string:
		return strconv.Atoi(v)
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, fmt.Errorf("ToInt: unknown supported type %T", v)
	}
}

func ToBool(v interface{}) (bool, error) {
	switch v := v.(type) {
	case bool:
		return v, nil
	case string:
		return strconv.ParseBool(v)
	case float64:
		return v != 0, nil
	case uint64:
		return v != 0, nil
	case int64:
		return v != 0, nil
	case int:
		return v != 0, nil
	default:
		return false, fmt.Errorf("ToInt: unknown supported type %T", v)
	}
}

// ToGeoPoint geo_point string with the format: "lat,lon".
func ToGeoPoint(v interface{}) ([]float64, error) {
	geoStr, err := ToString(v)
	if err != nil {
		return nil, err
	}
	geoStrSlice := strings.Split(geoStr, ",")
	if len(geoStrSlice) != 2 {
		return nil, fmt.Errorf("geo_point str must be lat,lon format")
	}

	geoPoint := make([]float64, 2)
	lat, err := strconv.ParseFloat(geoStrSlice[0], 64)
	if err != nil {
		return nil, fmt.Errorf("geo_point lat can't convert float64 type,value %T", geoStrSlice[0])
	}
	if !(lat >= -90 && lat <= 90.0) {
		return nil, fmt.Errorf("geo_point lat must between -90.0 and 90.0,vale %f", lat)
	}
	geoPoint[0] = lat

	lon, err := strconv.ParseFloat(geoStrSlice[1], 64)
	if err != nil {
		return nil, fmt.Errorf("geo_point lon can't convert float64 type,value %T", geoStrSlice[1])
	}
	if !(lon >= -180.0 && lat <= 180.0) {
		return nil, fmt.Errorf("geo_point lat must between -180.0 and 180.0,vale %f", lon)
	}
	geoPoint[1] = lon
	return geoPoint, nil
}
