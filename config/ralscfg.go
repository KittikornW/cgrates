/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package config

import (
	"time"

	"github.com/cgrates/cgrates/utils"
)

// Rater config section
type RalsCfg struct {
	RALsEnabled             bool            // start standalone server (no balancer)
	RALsThresholdSConns     []*HaPoolConfig // address where to reach ThresholdS config
	RALsStatSConns          []*HaPoolConfig
	RpSubjectPrefixMatching bool // enables prefix matching for the rating profile subject
	RemoveExpired           bool
	RALsMaxComputedUsage    map[string]time.Duration
}

//loadFromJsonCfg loads Rals config from JsonCfg
func (ralsCfg *RalsCfg) loadFromJsonCfg(jsnRALsCfg *RalsJsonCfg) (err error) {
	if jsnRALsCfg == nil {
		return nil
	}
	if jsnRALsCfg.Enabled != nil {
		ralsCfg.RALsEnabled = *jsnRALsCfg.Enabled
	}
	if jsnRALsCfg.Thresholds_conns != nil {
		ralsCfg.RALsThresholdSConns = make([]*HaPoolConfig, len(*jsnRALsCfg.Thresholds_conns))
		for idx, jsnHaCfg := range *jsnRALsCfg.Thresholds_conns {
			ralsCfg.RALsThresholdSConns[idx] = NewDfltHaPoolConfig()
			ralsCfg.RALsThresholdSConns[idx].loadFromJsonCfg(jsnHaCfg)
		}
	}
	if jsnRALsCfg.Stats_conns != nil {
		ralsCfg.RALsStatSConns = make([]*HaPoolConfig, len(*jsnRALsCfg.Stats_conns))
		for idx, jsnHaCfg := range *jsnRALsCfg.Stats_conns {
			ralsCfg.RALsStatSConns[idx] = NewDfltHaPoolConfig()
			ralsCfg.RALsStatSConns[idx].loadFromJsonCfg(jsnHaCfg)
		}
	}
	if jsnRALsCfg.Rp_subject_prefix_matching != nil {
		ralsCfg.RpSubjectPrefixMatching = *jsnRALsCfg.Rp_subject_prefix_matching
	}
	if jsnRALsCfg.Remove_expired != nil {
		ralsCfg.RemoveExpired = *jsnRALsCfg.Remove_expired
	}
	if jsnRALsCfg.Max_computed_usage != nil {
		for k, v := range *jsnRALsCfg.Max_computed_usage {
			if ralsCfg.RALsMaxComputedUsage[k], err = utils.ParseDurationWithNanosecs(v); err != nil {
				return
			}
		}
	}

	return nil
}
