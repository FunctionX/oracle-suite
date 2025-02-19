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

package sliceutil

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	m := []string{"a", "b", "c"}
	assert.Equal(t, m, Copy(m))
	assert.NotSame(t, m, Copy(m))
}

func TestMap(t *testing.T) {
	m := []string{"a", "b", "c"}
	assert.Equal(t, []string{"A", "B", "C"}, Map(m, strings.ToUpper))
	assert.NotSame(t, m, Map(m, strings.ToUpper))
}

func TestIsUnique(t *testing.T) {
	assert.True(t, IsUnique([]string{"a", "b", "c"}))
	assert.False(t, IsUnique([]string{"a", "b", "a"}))
}
