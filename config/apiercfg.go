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

// ApierCfg is the configuration of Apier service
type ApierCfg struct {
	CachesConns []*HaPoolConfig // connections towards Cache
}

func (aCfg *ApierCfg) loadFromJsonCfg(jsnCfg *ApierJsonCfg) (err error) {
	if jsnCfg == nil {
		return
	}
	if jsnCfg.Caches_conns != nil {
		aCfg.CachesConns = make([]*HaPoolConfig, len(*jsnCfg.Caches_conns))
		for idx, jsnHaCfg := range *jsnCfg.Caches_conns {
			aCfg.CachesConns[idx] = NewDfltHaPoolConfig()
			aCfg.CachesConns[idx].loadFromJsonCfg(jsnHaCfg)
		}
	}
	return nil
}
