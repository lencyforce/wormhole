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

package models

type Article struct {
	BaseModel

	UserId   uint   `gorm:"index" json:"-"`
	Title    string `gorm:"type:text" form:"title" json:"title" binding:"required"`
	Abstract string `gorm:"type:text" json:"abstract"`
	Content  string `gorm:"type:longtext" form:"content" json:"content" binding:"required"`
	Language string `gorm:"column:lang;size:64" json:"language"`

	ContentId  string `gorm:"unique_index" json:"content_id"`
	ContentDNA string `gorm:"unique_index" json:"content_dna"`
}

func (article *Article) DetectLanguage() string {
	return ""
}