//  Copyright (C) 2020 Maker Ecosystem Growth Holdings, INC.
//
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU Affero General Public License as
//  published by the Free Software Foundation, either version 3 of the
//  License, or (at your option) any later version.
//
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU Affero General Public License for more details.
//
//  You should have received a copy of the GNU Affero General Public License
//  along with this program.  If not, see <http://www.gnu.org/licenses/>.

package web

import (
	"net/http"
	"net/url"

	"github.com/makerdao/gofer/internal/marshal"
	"github.com/makerdao/gofer/pkg/cli"
	"github.com/makerdao/gofer/pkg/graph"
)

func PricesHandler(l graph.PriceModels) http.HandlerFunc {
	return marshallerHandler(func(m marshal.Marshaller, r *http.Request) error {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			return err
		}

		return cli.Prices(values["pair"], l, m)
	})
}
