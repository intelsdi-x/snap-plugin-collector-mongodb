/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mongo

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Mongo is an interface of Mongo calls
type Mongo interface {
	GetConnection(URL, AuthPassword, AuthUserName string) (*mgo.Session, error)
	GetServerStatus(database string, session *mgo.Session) (ServerStatus, error)
}

// GetConnection returns session mgo Object
func GetConnection(URL, AuthPassword, AuthUserName string) (*mgo.Session, error) {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{URL},
		Timeout:  1 * time.Second,
		Username: AuthUserName,
		Password: AuthPassword,
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}

// GetServerStatus returns server status
func GetServerStatus(session *mgo.Session) (ServerStatus, error) {
	result := ServerStatus{}
	if err := session.Run(bson.D{{"serverStatus", 1}}, &result); err != nil {
		return result, err
	}
	defer session.Close()

	return result, nil
}
