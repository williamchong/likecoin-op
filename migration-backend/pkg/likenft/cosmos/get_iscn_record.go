package cosmos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/util/httputil"
)

var ErrGetISCNRecord = errors.New("err getting iscn record")

func (a *LikeNFTCosmosClient) GetISCNRecord(
	iscnIdPrefix string,
	iscnVersionAtMint string,
) (*model.ISCN, error) {
	u, err := url.Parse("/iscn/records/id")

	if err != nil {
		return nil, errors.Join(ErrGetISCNRecord, fmt.Errorf("url.Parse %s", "/iscn/records/id"), err)
	}

	if iscnVersionAtMint == "" {
		iscnVersionAtMint = "1"
	}

	queryString := u.Query()
	queryString.Set("iscn_id", iscnIdPrefix)
	queryString.Set("iscn_version_at_mint", iscnVersionAtMint)
	u.RawQuery = queryString.Encode()

	base, err := url.Parse(a.NodeURL)
	if err != nil {
		return nil, errors.Join(ErrGetISCNRecord, fmt.Errorf("url.Parse %s", a.NodeURL), err)
	}

	resp, err := a.HTTPClient.Get(base.ResolveReference(u).String())

	if err != nil {
		return nil, errors.Join(ErrGetISCNRecord, fmt.Errorf("a.HTTPClient.Get"), err)
	}
	if err = httputil.HandleResponseStatus(resp); err != nil {
		return nil, errors.Join(ErrGetISCNRecord, err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var response model.ISCN
	err = decoder.Decode(&response)
	if err != nil {
		return nil, errors.Join(ErrGetISCNRecord, fmt.Errorf("decoder.Decode"), err)
	}
	return &response, nil
}
