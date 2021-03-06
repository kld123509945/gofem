// Copyright 2016 The Gofem Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shp

import (
	"math"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/utl"
)

// Ipoint implements integration point data: natural coordinates and weight
type Ipoint []float64 // len==4 => [r, s, t, w]

// ipsfactory holds all integration points sets
var ipsfactory = make(map[string][]Ipoint)

// GetIps returns a set of integration points
//  If the number (nips) of integration points is zero, it returns a default set
func (o *Shape) GetIps(nips, nipf int) (ips, ipf []Ipoint, err error) {

	// NURBS
	if o.Type == "nurbs" {
		maxord := o.Nurbs.Ord(0)
		for i := 1; i < o.Gndim; i++ {
			maxord = utl.Imax(maxord, o.Nurbs.Ord(i))
		}
		switch o.Gndim {
		case 1:
			switch maxord {
			case 1:
				ips = ips_lin_2
			case 2:
				ips = ips_lin_3
			default:
				err = chk.Err("1D NURBS: cannot get integration points with maxord=%d", maxord)
			}
		case 2:
			switch maxord {
			case 1:
				ips = ips_qua_4
				ipf = ips_lin_2
			case 2:
				ips = ips_qua_9
				ipf = ips_lin_3
			default:
				err = chk.Err("2D NURBS: cannot get integration points with maxord=%d", maxord)
			}
		case 3:
			switch maxord {
			case 1:
				ips = ips_hex_8
				ipf = ips_qua_4
			case 2:
				ips = ips_hex_27
				ipf = ips_qua_9
			default:
				err = chk.Err("3D NURBS: cannot get integration points with maxord=%d", maxord)
			}
		}
		return
	}

	// Lagrangean elements
	var ok bool
	key := io.Sf("%s_%d", o.Type, nips)
	ips, ok = ipsfactory[key]
	if !ok {
		err = chk.Err("cannot find integration point set for geometry type = %s and nips = %d\n", o.Type, nips)
		return
	}
	if o.Gndim > 1 {
		key = io.Sf("%s_%d", o.FaceType, nipf)
		ipf, ok = ipsfactory[key]
		if !ok {
			err = chk.Err("cannot find (face) integration point set for geometry type = %s and nips = %d\n", o.FaceType, nipf)
			return
		}
	}
	return
}

var (
	ips_lin_2 = []Ipoint{
		Ipoint{-math.Sqrt(3.0) / 3.0, 0.0, 0.0, 1.0},
		Ipoint{math.Sqrt(3.0) / 3.0, 0.0, 0.0, 1.0},
	}

	ips_lin_3 = []Ipoint{
		Ipoint{-math.Sqrt(3.0 / 5.0), 0.0, 0.0, 5.0 / 9.0},
		Ipoint{0.0, 0.0, 0.0, 8.0 / 9.0},
		Ipoint{math.Sqrt(3.0 / 5.0), 0.0, 0.0, 5.0 / 9.0},
	}

	ips_lin_5 = []Ipoint{
		Ipoint{-0.906179845938663992797627e+00, 0.0, 0.0, 0.23692688505618908751426e+00},
		Ipoint{-0.538469310105683091036314e+00, 0.0, 0.0, 0.47862867049936646804129e+00},
		Ipoint{0.000000000000000000000000e+00, 0.0, 0.0, 0.56888888888888888888889e+00},
		Ipoint{0.538469310105683091036314e+00, 0.0, 0.0, 0.47862867049936646804129e+00},
		Ipoint{0.906179845938663992797627e+00, 0.0, 0.0, 0.23692688505618908751426e+00},
	}

	ips_tri_1 = []Ipoint{
		Ipoint{1.0 / 3.0, 1.0 / 3.0, 0.0, 1.0 / 2.0},
	}

	ips_tri_3 = []Ipoint{
		Ipoint{1.0 / 6.0, 1.0 / 6.0, 0.0, 1.0 / 6.0},
		Ipoint{2.0 / 3.0, 1.0 / 6.0, 0.0, 1.0 / 6.0},
		Ipoint{1.0 / 6.0, 2.0 / 3.0, 0.0, 1.0 / 6.0},
	}

	ips_tri_12 = []Ipoint{
		Ipoint{0.873821971016996, 0.063089014491502, 0, 0.0254224531851035},
		Ipoint{0.063089014491502, 0.873821971016996, 0, 0.0254224531851035},
		Ipoint{0.063089014491502, 0.063089014491502, 0, 0.0254224531851035},
		Ipoint{0.501426509658179, 0.249286745170910, 0, 0.0583931378631895},
		Ipoint{0.249286745170910, 0.501426509658179, 0, 0.0583931378631895},
		Ipoint{0.249286745170910, 0.249286745170910, 0, 0.0583931378631895},
		Ipoint{0.053145049844817, 0.310352451033784, 0, 0.041425537809187},
		Ipoint{0.310352451033784, 0.053145049844817, 0, 0.041425537809187},
		Ipoint{0.053145049844817, 0.636502499121398, 0, 0.041425537809187},
		Ipoint{0.310352451033784, 0.636502499121398, 0, 0.041425537809187},
		Ipoint{0.636502499121398, 0.053145049844817, 0, 0.041425537809187},
		Ipoint{0.636502499121398, 0.310352451033784, 0, 0.041425537809187},
	}

	ips_tri_16 = []Ipoint{
		Ipoint{3.33333333333333E-01, 3.33333333333333E-01, 0.0, 7.21578038388935E-02},
		Ipoint{8.14148234145540E-02, 4.59292588292723E-01, 0.0, 4.75458171336425E-02},
		Ipoint{4.59292588292723E-01, 8.14148234145540E-02, 0.0, 4.75458171336425E-02},
		Ipoint{4.59292588292723E-01, 4.59292588292723E-01, 0.0, 4.75458171336425E-02},
		Ipoint{6.58861384496480E-01, 1.70569307751760E-01, 0.0, 5.16086852673590E-02},
		Ipoint{1.70569307751760E-01, 6.58861384496480E-01, 0.0, 5.16086852673590E-02},
		Ipoint{1.70569307751760E-01, 1.70569307751760E-01, 0.0, 5.16086852673590E-02},
		Ipoint{8.98905543365938E-01, 5.05472283170310E-02, 0.0, 1.62292488115990E-02},
		Ipoint{5.05472283170310E-02, 8.98905543365938E-01, 0.0, 1.62292488115990E-02},
		Ipoint{5.05472283170310E-02, 5.05472283170310E-02, 0.0, 1.62292488115990E-02},
		Ipoint{8.39477740995800E-03, 2.63112829634638E-01, 0.0, 1.36151570872175E-02},
		Ipoint{7.28492392955404E-01, 8.39477740995800E-03, 0.0, 1.36151570872175E-02},
		Ipoint{2.63112829634638E-01, 7.28492392955404E-01, 0.0, 1.36151570872175E-02},
		Ipoint{8.39477740995800E-03, 7.28492392955404E-01, 0.0, 1.36151570872175E-02},
		Ipoint{7.28492392955404E-01, 2.63112829634638E-01, 0.0, 1.36151570872175E-02},
		Ipoint{2.63112829634638E-01, 8.39477740995800E-03, 0.0, 1.36151570872175E-02},
	}

	ips_qua_4 = []Ipoint{
		Ipoint{-math.Sqrt(3.0) / 3.0, -math.Sqrt(3.0) / 3.0, 0.0, 1.0},
		Ipoint{math.Sqrt(3.0) / 3.0, -math.Sqrt(3.0) / 3.0, 0.0, 1.0},
		Ipoint{-math.Sqrt(3.0) / 3.0, math.Sqrt(3.0) / 3.0, 0.0, 1.0},
		Ipoint{math.Sqrt(3.0) / 3.0, math.Sqrt(3.0) / 3.0, 0.0, 1.0},
	}

	ips_qua_9 = []Ipoint{
		Ipoint{-math.Sqrt(3.0 / 5.0), -math.Sqrt(3.0 / 5.0), 0.0, 25.0 / 81.0},
		Ipoint{0.0, -math.Sqrt(3.0 / 5.0), 0.0, 40.0 / 81.0},
		Ipoint{math.Sqrt(3.0 / 5.0), -math.Sqrt(3.0 / 5.0), 0.0, 25.0 / 81.0},
		Ipoint{-math.Sqrt(3.0 / 5.0), 0.0, 0.0, 40.0 / 81.0},
		Ipoint{0.0, 0.0, 0.0, 64.0 / 81.0},
		Ipoint{math.Sqrt(3.0 / 5.0), 0.0, 0.0, 40.0 / 81.0},
		Ipoint{-math.Sqrt(3.0 / 5.0), math.Sqrt(3.0 / 5.0), 0.0, 25.0 / 81.0},
		Ipoint{0.0, math.Sqrt(3.0 / 5.0), 0.0, 40.0 / 81.0},
		Ipoint{math.Sqrt(3.0 / 5.0), math.Sqrt(3.0 / 5.0), 0.0, 25.0 / 81.0},
	}

	ips_tet_1 = []Ipoint{
		Ipoint{1.0 / 4.0, 1.0 / 4.0, 1.0 / 4.0, 1.0 / 6.0},
	}

	ips_tet_4 = []Ipoint{
		Ipoint{(5.0 + 3.0*math.Sqrt(5.0)) / 20.0, (5.0 - math.Sqrt(5.0)) / 20.0, (5.0 - math.Sqrt(5.0)) / 20.0, 1.0 / 24},
		Ipoint{(5.0 - math.Sqrt(5.0)) / 20.0, (5.0 + 3.0*math.Sqrt(5.0)) / 20.0, (5.0 - math.Sqrt(5.0)) / 20.0, 1.0 / 24},
		Ipoint{(5.0 - math.Sqrt(5.0)) / 20.0, (5.0 - math.Sqrt(5.0)) / 20.0, (5.0 + 3.0*math.Sqrt(5.0)) / 20.0, 1.0 / 24},
		Ipoint{(5.0 - math.Sqrt(5.0)) / 20.0, (5.0 - math.Sqrt(5.0)) / 20.0, (5.0 - math.Sqrt(5.0)) / 20.0, 1.0 / 24},
	}

	ips_tet_5 = []Ipoint{
		Ipoint{1.0 / 4.0, 1.0 / 4.0, 1.0 / 4.0, -2.0 / 15.0},
		Ipoint{1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0, 3.0 / 40.0},
		Ipoint{1.0 / 6.0, 1.0 / 6.0, 1.0 / 2.0, 3.0 / 40.0},
		Ipoint{1.0 / 6.0, 1.0 / 2.0, 1.0 / 6.0, 3.0 / 40.0},
		Ipoint{1.0 / 2.0, 1.0 / 6.0, 1.0 / 6.0, 3.0 / 40.0},
	}

	ips_hex_6 = []Ipoint{
		Ipoint{1.0, 0.0, 0.0, 4.0 / 3.0},
		Ipoint{-1.0, 0.0, 0.0, 4.0 / 3.0},
		Ipoint{0.0, 1.0, 0.0, 4.0 / 3.0},
		Ipoint{0.0, -1.0, 0.0, 4.0 / 3.0},
		Ipoint{0.0, 0.0, 1.0, 4.0 / 3.0},
		Ipoint{0.0, 0.0, -1.0, 4.0 / 3.0},
	}

	ips_hex_8 = []Ipoint{
		Ipoint{-math.Sqrt(3.0) / 3.0, -math.Sqrt(3.0) / 3.0, -math.Sqrt(3.0) / 3.0, 1.0},
		Ipoint{math.Sqrt(3.0) / 3.0, -math.Sqrt(3.0) / 3.0, -math.Sqrt(3.0) / 3.0, 1.0},
		Ipoint{-math.Sqrt(3.0) / 3.0, math.Sqrt(3.0) / 3.0, -math.Sqrt(3.0) / 3.0, 1.0},
		Ipoint{math.Sqrt(3.0) / 3.0, math.Sqrt(3.0) / 3.0, -math.Sqrt(3.0) / 3.0, 1.0},
		Ipoint{-math.Sqrt(3.0) / 3.0, -math.Sqrt(3.0) / 3.0, math.Sqrt(3.0) / 3.0, 1.0},
		Ipoint{math.Sqrt(3.0) / 3.0, -math.Sqrt(3.0) / 3.0, math.Sqrt(3.0) / 3.0, 1.0},
		Ipoint{-math.Sqrt(3.0) / 3.0, math.Sqrt(3.0) / 3.0, math.Sqrt(3.0) / 3.0, 1.0},
		Ipoint{math.Sqrt(3.0) / 3.0, math.Sqrt(3.0) / 3.0, math.Sqrt(3.0) / 3.0, 1.0},
	}

	ips_hex_14 = []Ipoint{
		Ipoint{math.Sqrt(19.0 / 30.0), 0.0, 0.0, 320.0 / 361.0},
		Ipoint{-math.Sqrt(19.0 / 30.0), 0.0, 0.0, 320.0 / 361.0},
		Ipoint{0.0, math.Sqrt(19.0 / 30.0), 0.0, 320.0 / 361.0},
		Ipoint{0.0, -math.Sqrt(19.0 / 30.0), 0.0, 320.0 / 361.0},
		Ipoint{0.0, 0.0, math.Sqrt(19.0 / 30.0), 320.0 / 361.0},
		Ipoint{0.0, 0.0, -math.Sqrt(19.0 / 30.0), 320.0 / 361.0},
		Ipoint{math.Sqrt(19.0 / 33.0), math.Sqrt(19.0 / 33.0), math.Sqrt(19.0 / 33.0), 121.0 / 361.0},
		Ipoint{-math.Sqrt(19.0 / 33.0), math.Sqrt(19.0 / 33.0), math.Sqrt(19.0 / 33.0), 121.0 / 361.0},
		Ipoint{math.Sqrt(19.0 / 33.0), -math.Sqrt(19.0 / 33.0), math.Sqrt(19.0 / 33.0), 121.0 / 361.0},
		Ipoint{-math.Sqrt(19.0 / 33.0), -math.Sqrt(19.0 / 33.0), math.Sqrt(19.0 / 33.0), 121.0 / 361.0},
		Ipoint{math.Sqrt(19.0 / 33.0), math.Sqrt(19.0 / 33.0), -math.Sqrt(19.0 / 33.0), 121.0 / 361.0},
		Ipoint{-math.Sqrt(19.0 / 33.0), math.Sqrt(19.0 / 33.0), -math.Sqrt(19.0 / 33.0), 121.0 / 361.0},
		Ipoint{math.Sqrt(19.0 / 33.0), -math.Sqrt(19.0 / 33.0), -math.Sqrt(19.0 / 33.0), 121.0 / 361.0},
		Ipoint{-math.Sqrt(19.0 / 33.0), -math.Sqrt(19.0 / 33.0), -math.Sqrt(19.0 / 33.0), 121.0 / 361.0},
	}

	ips_hex_27 = []Ipoint{
		Ipoint{-0.774596669241483, -0.774596669241483, -0.774596669241483, 0.171467764060357},
		Ipoint{0.000000000000000, -0.774596669241483, -0.774596669241483, 0.274348422496571},
		Ipoint{0.774596669241483, -0.774596669241483, -0.774596669241483, 0.171467764060357},
		Ipoint{-0.774596669241483, 0.000000000000000, -0.774596669241483, 0.274348422496571},
		Ipoint{0.000000000000000, 0.000000000000000, -0.774596669241483, 0.438957475994513},
		Ipoint{0.774596669241483, 0.000000000000000, -0.774596669241483, 0.274348422496571},
		Ipoint{-0.774596669241483, 0.774596669241483, -0.774596669241483, 0.171467764060357},
		Ipoint{0.000000000000000, 0.774596669241483, -0.774596669241483, 0.274348422496571},
		Ipoint{0.774596669241483, 0.774596669241483, -0.774596669241483, 0.171467764060357},
		Ipoint{-0.774596669241483, -0.774596669241483, 0.000000000000000, 0.274348422496571},
		Ipoint{0.000000000000000, -0.774596669241483, 0.000000000000000, 0.438957475994513},
		Ipoint{0.774596669241483, -0.774596669241483, 0.000000000000000, 0.274348422496571},
		Ipoint{-0.774596669241483, 0.000000000000000, 0.000000000000000, 0.438957475994513},
		Ipoint{0.000000000000000, 0.000000000000000, 0.000000000000000, 0.702331961591221},
		Ipoint{0.774596669241483, 0.000000000000000, 0.000000000000000, 0.438957475994513},
		Ipoint{-0.774596669241483, 0.774596669241483, 0.000000000000000, 0.274348422496571},
		Ipoint{0.000000000000000, 0.774596669241483, 0.000000000000000, 0.438957475994513},
		Ipoint{0.774596669241483, 0.774596669241483, 0.000000000000000, 0.274348422496571},
		Ipoint{-0.774596669241483, -0.774596669241483, 0.774596669241483, 0.171467764060357},
		Ipoint{0.000000000000000, -0.774596669241483, 0.774596669241483, 0.274348422496571},
		Ipoint{0.774596669241483, -0.774596669241483, 0.774596669241483, 0.171467764060357},
		Ipoint{-0.774596669241483, 0.000000000000000, 0.774596669241483, 0.274348422496571},
		Ipoint{0.000000000000000, 0.000000000000000, 0.774596669241483, 0.438957475994513},
		Ipoint{0.774596669241483, 0.000000000000000, 0.774596669241483, 0.274348422496571},
		Ipoint{-0.774596669241483, 0.774596669241483, 0.774596669241483, 0.171467764060357},
		Ipoint{0.000000000000000, 0.774596669241483, 0.774596669241483, 0.274348422496571},
		Ipoint{0.774596669241483, 0.774596669241483, 0.774596669241483, 0.171467764060357},
	}
)
