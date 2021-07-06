// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aliyun

// https://help.aliyun.com/document_detail/132255.html
var esSpec = map[string]struct {
	VcpuCount  int
	VmemSizeGb int
}{
	"elasticsearch.ic5.2xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 8, VmemSizeGb: 8},
	"elasticsearch.ic5.3xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 12, VmemSizeGb: 12},
	"elasticsearch.ic5.4xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 16, VmemSizeGb: 16},
	"elasticsearch.ic5.large": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 16, VmemSizeGb: 16},
	"elasticsearch.ic5.xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 4, VmemSizeGb: 4},
	"elasticsearch.n4.small": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 1, VmemSizeGb: 2},
	"elasticsearch.sn1ne.2xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 8, VmemSizeGb: 16},
	"elasticsearch.sn1ne.4xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 16, VmemSizeGb: 32},
	"elasticsearch.sn1ne.8xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 32, VmemSizeGb: 64},
	"elasticsearch.sn1ne.large": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 2, VmemSizeGb: 4},
	"elasticsearch.sn1ne.xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 4, VmemSizeGb: 8},
	"elasticsearch.sn2ne.2xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 8, VmemSizeGb: 32},
	"elasticsearch.sn2ne.4xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 16, VmemSizeGb: 64},
	"elasticsearch.sn2ne.8xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 32, VmemSizeGb: 128},
	"elasticsearch.sn2ne.large": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 2, VmemSizeGb: 8},
	"elasticsearch.sn2ne.xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 4, VmemSizeGb: 16},
	"elasticsearch.r6.8xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 32, VmemSizeGb: 256},
	"elasticsearch.r6.4xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 16, VmemSizeGb: 128},
	"elasticsearch.r5.2xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 8, VmemSizeGb: 64},
	"elasticsearch.r5.large": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 2, VmemSizeGb: 16},
	"elasticsearch.r5.xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 4, VmemSizeGb: 32},
	"elasticsearch.d1.2xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 8, VmemSizeGb: 32},
	"elasticsearch.d1.4xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 16, VmemSizeGb: 64},
	"elasticsearch.i2g.2xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 8, VmemSizeGb: 32},
	"elasticsearch.i2g.4xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 16, VmemSizeGb: 64},
	"elasticsearch.i2g.8xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 32, VmemSizeGb: 128},
	"elasticsearch.d2s.5xlarge": struct {
		VcpuCount  int
		VmemSizeGb int
	}{VcpuCount: 20, VmemSizeGb: 88},
}
