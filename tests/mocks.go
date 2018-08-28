/*
 * Copyright 2018 Primas Lab Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tests

import (
	"github.com/primasio/wormhole/models"
	"github.com/primasio/wormhole/util"
	"log"
)

func CreateTestUser() (user *models.User, err error) {

	u := &models.User{}

	randStr := util.RandString(5)

	u.Username = "test_user_" + randStr
	u.Nickname = "Test User " + randStr
	u.Password = "PrimasGoGoGo"

	log.Println("Created test user: " + u.Username)

	return u, nil
}
