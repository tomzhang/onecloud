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

package models

import (
	compute_models "yunion.io/x/onecloud/pkg/compute/models"
)

type Vpc struct {
	compute_models.SVpc

	Networks Networks `json:"-"`
}

func (el *Vpc) Copy() *Vpc {
	return &Vpc{
		SVpc: el.SVpc,
	}
}

type Network struct {
	compute_models.SNetwork
	// returned as extra column
	VpcId string

	Vpc           *Vpc          `json:"-"`
	Guestnetworks Guestnetworks `json:"-"`
}

func (el *Network) Copy() *Network {
	return &Network{
		SNetwork: el.SNetwork,
		VpcId:    el.VpcId,
	}
}

type Guestnetwork struct {
	compute_models.SGuestnetwork

	Guest   *Guest   `json:"-"`
	Network *Network `json:"-"`
}

func (el *Guestnetwork) Copy() *Guestnetwork {
	return &Guestnetwork{
		SGuestnetwork: el.SGuestnetwork,
	}
}

type Guest struct {
	compute_models.SGuest

	Host               *Host          `json:"-"`
	AdminSecurityGroup *SecurityGroup `json:"-"`
	SecurityGroups     SecurityGroups `json:"-"`
}

func (el *Guest) Copy() *Guest {
	return &Guest{
		SGuest: el.SGuest,
	}
}

type Host struct {
	compute_models.SHost
}

func (el *Host) Copy() *Host {
	return &Host{
		SHost: el.SHost,
	}
}

type Guestsecgroup struct {
	compute_models.SGuestsecgroup

	Guest         *Guest         `json:"-"`
	SecurityGroup *SecurityGroup `json:"-"`
}

func (el *Guestsecgroup) ModelSetKey() string {
	return el.GuestId + "/" + el.SecgroupId
}

func (el *Guestsecgroup) Copy() *Guestsecgroup {
	return &Guestsecgroup{
		SGuestsecgroup: el.SGuestsecgroup,
	}
}

type SecurityGroup struct {
	compute_models.SSecurityGroup

	SecurityGroupRules SecurityGroupRules `json:"-"`
}

func (el *SecurityGroup) Copy() *SecurityGroup {
	return &SecurityGroup{
		SSecurityGroup: el.SSecurityGroup,
	}
}

type SecurityGroupRule struct {
	compute_models.SSecurityGroupRule

	SecurityGroup *SecurityGroup `json:"-"`
}

func (el *SecurityGroupRule) Copy() *SecurityGroupRule {
	return &SecurityGroupRule{
		SSecurityGroupRule: el.SSecurityGroupRule,
	}
}
