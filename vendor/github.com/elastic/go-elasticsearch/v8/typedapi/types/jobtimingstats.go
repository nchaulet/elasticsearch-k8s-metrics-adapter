// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// JobTimingStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/Job.ts#L109-L118
type JobTimingStats struct {
	AverageBucketProcessingTimeMs                   *float64 `json:"average_bucket_processing_time_ms,omitempty"`
	BucketCount                                     int64    `json:"bucket_count"`
	ExponentialAverageBucketProcessingTimeMs        *float64 `json:"exponential_average_bucket_processing_time_ms,omitempty"`
	ExponentialAverageBucketProcessingTimePerHourMs float64  `json:"exponential_average_bucket_processing_time_per_hour_ms"`
	JobId                                           string   `json:"job_id"`
	MaximumBucketProcessingTimeMs                   *float64 `json:"maximum_bucket_processing_time_ms,omitempty"`
	MinimumBucketProcessingTimeMs                   *float64 `json:"minimum_bucket_processing_time_ms,omitempty"`
	TotalBucketProcessingTimeMs                     float64  `json:"total_bucket_processing_time_ms"`
}

// NewJobTimingStats returns a JobTimingStats.
func NewJobTimingStats() *JobTimingStats {
	r := &JobTimingStats{}

	return r
}
