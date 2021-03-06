/*
 * Copyright (C) 2018 The Seer Authors. All rights reserved.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package kalman_test

import (
	"testing"

	"github.com/cshenton/seer/kalman"
	"gonum.org/v1/gonum/mat"
)

func TestNewState(t *testing.T) {
	tt := []struct {
		name   string
		loc    *mat.Dense
		cov    *mat.Dense
		errNil bool
	}{
		{"Non column loc", mat.NewDense(1, 2, []float64{0, 0}), mat.NewDense(1, 1, []float64{1}), false},
		{"Non square cov", mat.NewDense(1, 1, []float64{0}), mat.NewDense(1, 2, []float64{1, 1}), false},
		{"Non matching loc, cov", mat.NewDense(2, 1, []float64{0, 0}), mat.NewDense(1, 1, []float64{1}), false},
		{"Valid dim 1", mat.NewDense(1, 1, []float64{0}), mat.NewDense(1, 1, []float64{1}), true},
		{"Valid dim 2", mat.NewDense(2, 1, []float64{0, 0}), mat.NewDense(2, 2, []float64{1, 0, 0, 1}), true},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := kalman.NewState(tc.loc, tc.cov)
			errNil := (err == nil)
			if errNil != tc.errNil {
				t.Errorf("Expected error == nil to be %v, but it was %v", tc.errNil, errNil)
			}
		})
	}
}

func TestStateDim(t *testing.T) {
	tt := []struct {
		name string
		loc  *mat.Dense
		cov  *mat.Dense
		dim  int
	}{
		{"Dim 1", mat.NewDense(1, 1, []float64{0}), mat.NewDense(1, 1, []float64{1}), 1},
		{"Dim 2", mat.NewDense(2, 1, []float64{0, 0}), mat.NewDense(2, 2, []float64{1, 0, 0, 1}), 2},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s, err := kalman.NewState(tc.loc, tc.cov)
			if err != nil {
				t.Fatal("Failed to construct state")
			}
			d := s.Dim()
			if tc.dim != d {
				t.Errorf("Expected Dim to be %v, but it was %v", tc.dim, d)
			}
		})
	}
}
