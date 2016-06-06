/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// Creates a CreateInstances struct based on a given service
func buildCreateInstances(s *service) InstancesCreate {
	messages := []MonitorMessage{}
	messages = append(messages, MonitorMessage{Body: "Creating instances:", Level: "INFO"})

	res := InstancesCreate{}

	if len(s.InstancesToCreate.Items) > 0 {
		res, messages = buildInstancesList(s, s.InstancesToCreate.Items, messages, false)
	} else {
		res, messages = buildInstancesList(s, s.Instances.Items, messages, false)
	}

	UserOutput(s.Channel(), messages)

	return res
}

// Creates an UpdateInstances struct based on a given service
func buildUpdateInstances(s *service) InstancesCreate {
	messages := []MonitorMessage{}
	messages = append(messages, MonitorMessage{Body: "Updating instances:", Level: "INFO"})

	res, messages := buildInstancesList(s, s.InstancesToUpdate.Items, messages, true)

	UserOutput(s.Channel(), messages)

	return res
}

// Creates a DeleteInstances struct based on a given service
func buildDeleteInstances(s *service) InstancesCreate {
	messages := []MonitorMessage{}
	messages = append(messages, MonitorMessage{Body: "Deleting instances:", Level: "INFO"})

	res, messages := buildInstancesList(s, s.InstancesToDelete.Items, messages, false)

	UserOutput(s.Channel(), messages)

	return res
}

func buildInstancesList(s *service, list []instance, messages []MonitorMessage, sequential bool) (InstancesCreate, []MonitorMessage) {
	d := s.datacenter()

	m := InstancesCreate{
		Service:              s.ID,
		Instances:            list,
		SequentialProcessing: sequential,
	}

	for i, ii := range list {
		messages = append(messages, MonitorMessage{Body: "\t - " + ii.Name, Level: ""})

		m.Instances[i] = instance{
			Name:               ii.Name,
			Type:               ii.Type,
			CPU:                ii.CPU,
			RAM:                ii.RAM,
			IP:                 ii.IP,
			Catalog:            ii.Catalog,
			Image:              ii.Image,
			Disks:              ii.Disks,
			NetworkName:        ii.NetworkName,
			ClientName:         s.ClientName,
			DatacenterName:     d.Name,
			DatacenterPassword: d.Password,
			DatacenterRegion:   d.Region,
			DatacenterType:     d.Type,
			DatacenterUsername: d.Username,
			VCloudURL:          d.VCloudURL,
		}
		m.Instances[i].Status = ii.Status
	}

	return m, messages
}