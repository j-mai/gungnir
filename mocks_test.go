/**
 * Copyright 2019 Comcast Cable Communications Management, LLC
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
 *
 */

package main

import (
	"github.com/stretchr/testify/mock"
	"github.com/xmidt-org/codex-db"
	"github.com/xmidt-org/voynicrypto"
)

type mockRecordGetter struct {
	mock.Mock
}

func (rg *mockRecordGetter) GetRecords(deviceID string, limit int, stateHash string) ([]db.Record, error) {
	args := rg.Called(deviceID, limit, stateHash)
	return args.Get(0).([]db.Record), args.Error(1)
}

func (rg *mockRecordGetter) GetRecordsOfType(deviceID string, limit int, eventType db.EventType, stateHash string) ([]db.Record, error) {
	args := rg.Called(deviceID, limit, eventType, stateHash)
	return args.Get(0).([]db.Record), args.Error(1)
}

func (rg *mockRecordGetter) GetStateHash(records []db.Record) (string, error) {
	args := rg.Called(records)
	return args.String(0), args.Error(1)
}

type mockDecrypter struct {
	mock.Mock
}

func (md *mockDecrypter) DecryptMessage(cipher []byte, nonce []byte) ([]byte, error) {
	args := md.Called(cipher, nonce)
	return cipher, args.Error(0)
}

func (*mockDecrypter) GetAlgorithm() voynicrypto.AlgorithmType {
	return voynicrypto.None
}

func (*mockDecrypter) GetKID() string {
	return "none"
}
