package odin

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/cybledev/odin-sdk-go/models"
	"github.com/cybledev/odin-sdk-go/odin/certificate"
	"github.com/cybledev/odin-sdk-go/odin/domain"
	"github.com/cybledev/odin-sdk-go/odin/exposed_buckets"
	"github.com/cybledev/odin-sdk-go/odin/exposed_files"
	"github.com/cybledev/odin-sdk-go/odin/hosts"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"
)

var (
	client = NewHTTPClient(nil)
	ctx    = context.Background()

	url        = "google.com"
	exploitsIP = "172.232.206.199"
)

func apiKeyAuth() runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam("x-api-key", os.Getenv("odin_api_key"))
	})
}

// list all files from the bucket
func TestPostExposedBucketFilesList(t *testing.T) {
	limit := int64(10)
	esr := models.ExposedSearchRequest{
		Limit: &limit,
		Query: "bucket:\"lit-link-prd.appspot.com\"",
	}

	searchParams := exposed_files.NewPostV1ExposedFilesSearchParamsWithContext(ctx).WithQuery(&esr)
	resp, err := client.ExposedFiles.PostV1ExposedFilesSearch(searchParams, apiKeyAuth())
	require.Nil(t, err)
	require.Equal(t, resp.IsSuccess(), true)
	require.NotNil(t, resp)

	t.Logf(resp.String())
}

func TestPostExposedBucketSearchPagination(t *testing.T) {
	limit := int64(10)
	esr := models.ExposedSearchRequest{
		Limit:   &limit,
		Query:   "provider: aws",
		SortBy:  "files",
		SortDir: "desc",
	}

	buckets := []*models.ExposedBucket{}
	searchParams := exposed_buckets.NewPostV1ExposedBucketsSearchParamsWithContext(ctx).WithQuery(&esr)
	for i := 0; i < 3; i++ {
		resp, err := client.ExposedBuckets.PostV1ExposedBucketsSearch(searchParams, apiKeyAuth())
		require.Nil(t, err)
		require.Equal(t, resp.IsSuccess(), true)
		require.NotNil(t, resp)

		esr.Start = resp.Payload.Pagination.Last
		buckets = append(buckets, resp.Payload.Data...)
	}

	bucketsMarshalled, err := json.Marshal(buckets)
	require.Nil(t, err)
	t.Logf(string(bucketsMarshalled))
}

func TestPostHostsCount(t *testing.T) {
	query := models.CountRequest{
		Query: "(last_updated_at:[\"2024-07-08T02:41:15.528Z\" TO *] AND services.port:80) OR asn.number:AS63949",
	}
	params := hosts.NewPostV1HostsCountParamsWithContext(ctx).WithQuery(&query)

	resp, err := client.Hosts.PostV1HostsCount(params, apiKeyAuth())
	require.Nil(t, err)
	require.Equal(t, resp.IsSuccess(), true)
	require.NotNil(t, resp)

	t.Logf(resp.String())
}

func TestPostCertificatesSearch(t *testing.T) {
	limit := int64(1)
	query := models.CertificateCertSearchRequest{
		Limit: &limit,
		Query: "certificate.subject_alt_name.dns_names:'cloudflare.com' AND certificate.validity.not_after:\"2024-09-20T18:19:24\"",
	}
	params := certificate.NewPostV1CertificatesSearchParamsWithContext(ctx).WithQuery(&query)
	resp, err := client.Certificate.PostV1CertificatesSearch(params, apiKeyAuth())

	require.Nil(t, err)
	require.Equal(t, resp.IsSuccess(), true)
	require.NotNil(t, resp)

	t.Logf(resp.String())
}

func TestPostHostsSearch(t *testing.T) {
	limit := int64(1)
	query := models.SearchRequest{
		Limit: &limit,
		Query: "services.port:80 AND asn.number:AS63949",
	}

	records := make([]*models.IpservicesIPSummaryData, 0)
	for i := 0; i < 3; i++ {
		params := hosts.NewPostV1HostsSearchParamsWithContext(ctx).WithQuery(&query)

		res, err := client.Hosts.PostV1HostsSearch(params, apiKeyAuth())
		require.Nil(t, err)

		records = append(records, res.Payload.Data...)
		query.Start = res.Payload.Pagination.Last
	}

	recordsMarshalled, err := json.Marshal(records)
	require.Nil(t, err)

	t.Log(string(recordsMarshalled))
}

func TestPostHostsV2Search(t *testing.T) {
	limit := int64(1)
	query := models.CybleComOdinAPIControllersV2IpservicesSearchRequest{
		Limit: &limit,
		Query: "asn.number:'AS13336'",
	}

	params := hosts.NewPostV2HostsSearchParamsWithContext(ctx).WithQuery(&query)
	res, err := client.Hosts.PostV2HostsSearch(params, apiKeyAuth())
	require.Nil(t, err)

	t.Log(res.String())
}

func TestPostExposedBucketSearch(t *testing.T) {
	limit := int64(10)
	esr := models.ExposedSearchRequest{
		Limit: &limit,
		Query: "test",
	}

	searchParams := exposed_buckets.NewPostV1ExposedBucketsSearchParamsWithContext(ctx).WithQuery(&esr)
	resp, err := client.ExposedBuckets.PostV1ExposedBucketsSearch(searchParams, apiKeyAuth())

	require.Nil(t, err)
	require.Equal(t, resp.IsSuccess(), true)
	require.NotNil(t, resp)

	t.Logf(resp.String())
}

func TestPostExposedFileSearch(t *testing.T) {
	limit := int64(10)
	esr := models.ExposedSearchRequest{
		Limit: &limit,
		Query: "test",
	}

	searchParams := exposed_files.NewPostV1ExposedFilesSearchParamsWithContext(ctx).WithQuery(&esr)
	resp, err := client.ExposedFiles.PostV1ExposedFilesSearch(searchParams, apiKeyAuth())
	require.Nil(t, err)
	require.Equal(t, resp.IsSuccess(), true)
	require.NotNil(t, resp)

	t.Logf(resp.String())
}

func TestPostHostsSummary(t *testing.T) {
	limit := int64(1)

	field := "services.port"
	query := &models.SummaryRequest{
		Limit: &limit,
		Field: &field,
	}
	params := hosts.NewPostV1HostsSummaryParamsWithContext(ctx).WithQuery(query)
	res, err := client.Hosts.PostV1HostsSummary(params, apiKeyAuth())
	require.Nil(t, err)
	require.Equal(t, res.IsSuccess(), true)
	require.NotNil(t, res)

	t.Log(res.String())
}

func TestHostsCveIP(t *testing.T) {
	params := hosts.NewGetV1HostsCveIPParamsWithContext(ctx).WithIP(exploitsIP)
	resp, err := client.Hosts.GetV1HostsCveIP(params, apiKeyAuth())
	require.Nil(t, err)
	require.Equal(t, resp.IsSuccess(), true)
	require.NotNil(t, resp)

	t.Log(resp.String())
}

func TestHostsExploitsIP(t *testing.T) {
	params := hosts.NewGetV1HostsExploitsIPParamsWithContext(ctx).WithIP(exploitsIP)
	resp, err := client.Hosts.GetV1HostsExploitsIP(params, apiKeyAuth())
	require.Nil(t, err)
	require.Equal(t, resp.IsSuccess(), true)
	require.NotNil(t, resp)

	t.Logf(resp.String())
}

func TestGetHostsIPParams(t *testing.T) {
	params := hosts.NewGetV1HostsIPParamsWithContext(ctx).WithIP(exploitsIP)
	client.Hosts.GetV1HostsIP(params, apiKeyAuth())
}

func TestPostCertificatesSummary(t *testing.T) {
	field := "tags"
	limit := int64(5)
	query := models.CertificateCertSummaryRequest{
		Field: &field,
		Limit: &limit,
	}

	params := certificate.NewPostV1CertificatesSummaryParamsWithContext(ctx).WithQuery(&query)
	resp, err := client.Certificate.PostV1CertificatesSummary(params, apiKeyAuth())
	require.Nil(t, err)
	require.Equal(t, resp.IsSuccess(), true)
	require.NotNil(t, resp)

	t.Logf(resp.String())
}

func TestPostCertificatesCount(t *testing.T) {
	query := &models.CertificateCertCountRequest{}
	params := certificate.NewPostV1CertificatesCountParamsWithContext(ctx).WithQuery(query)
	resp, err := client.Certificate.PostV1CertificatesCount(params, apiKeyAuth())
	require.Nil(t, err)
	require.Equal(t, resp.IsSuccess(), true)
	require.NotNil(t, resp)

	t.Logf(resp.String())
}

func TestGetCertificatesHash(t *testing.T) {
	hashVal := "DBE5E294A5D0D76E8B44AB86C623F51BCB8B247356A5253BFCCA5F425E8EBD09"

	params := certificate.NewGetV1CertificatesHashParamsWithContext(ctx).WithHash(hashVal)
	resp, err := client.Certificate.GetV1CertificatesHash(params, apiKeyAuth())

	require.Nil(t, err)
	require.Equal(t, resp.IsSuccess(), true)
	require.NotNil(t, resp)

	t.Logf(resp.String())
}

func TestPostDNSSearch(t *testing.T) {

	limit := int64(10)
	query := models.DNSDomainRequest{
		Domain: url,
		Limit:  &limit,
	}
	params := domain.NewPostV1DomainSearchParamsWithContext(ctx).WithQuery(&query)
	client.Domain.PostV1DomainSearch(params, apiKeyAuth())
}
