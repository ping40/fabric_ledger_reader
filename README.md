## build

- GO111MODULE=on
- go build -o ledgerReader

## run

- ./ledgerReader -file_name  /tmp/blockfile_000000  -number_end 35

result:

```json
{
  "header": {
    "number": 34,
    "previous_hash": "hwgIqkKsLO2g2Axp3GKZPgv7S9gzJfF4y43dAu/N7hU=",
    "data_hash": "AWYv5PfByFGS7Sy9v4lDFW0/2gBWeQFwg1Z7HcTt/TU="
  },
  "data": {
    "data": [
      {
        "payload": {
          "header": {
            "channel_header": {
              "type": 3,
              "timestamp": {
                "seconds": 1559549179,
                "nanos": 477864009
              },
              "channel_id": "mychannel",
              "tx_id": "f0e5b3021fcc80ce2c5ede04871ac00ca63202c6c14e009098e6afdfcfbe3e22",
              "extension": "EgYSBG15Y2M="
            },
            "signature_header": {
              "creator": {
                "mspid": "Org2MSP",
                "id_bytes": "NLKGwNPbQR6ZB1IXm52vbQ=="
              },
              "nonce": "6SdWUDuCQAb5NoYos0LuXGFNBZ3V7Erw"
            }
          },
          "endorser_transaction": {
            "actions": [
              {
                "header": {
                  "creator": {
                    "mspid": "Org2MSP",
                    "id_bytes": "NLKGwNPbQR6ZB1IXm52vbQ=="
                  },
                  "nonce": "6SdWUDuCQAb5NoYos0LuXGFNBZ3V7Erw"
                },
                "payload": {
                  "chaincode_proposal_payload": {
                    "proposal_hash": "Ch4IARIGEgRteWNjGhIKBmludm9rZQoBYQoBYgoCMTA=",
                    "extension": {}
                  },
                  "action": {
                    "proposal_response_payload": {
                      "proposal_hash": "EvBFpiT2wQ4g81jd2YHd4QkmX5+Yy3V5kql0ijvLdr8=",
                      "extension": {
                        "results": {
                          "NsRwSets": [
                            {
                              "NameSpace": "lscc",
                              "KvRwSet": {
                                "reads": [
                                  {
                                    "key": "mycc",
                                    "version": {
                                      "block_num": 3
                                    }
                                  }
                                ]
                              },
                              "CollHashedRwSets": null
                            },
                            {
                              "NameSpace": "mycc",
                              "KvRwSet": {
                                "reads": [
                                  {
                                    "key": "a",
                                    "version": {
                                      "block_num": 33
                                    }
                                  },
                                  {
                                    "key": "b",
                                    "version": {
                                      "block_num": 33
                                    }
                                  }
                                ],
                                "writes": [
                                  {
                                    "key": "a",
                                    "value": "NDc5MA=="
                                  },
                                  {
                                    "key": "b",
                                    "value": "NTUxMA=="
                                  }
                                ]
                              },
                              "CollHashedRwSets": null
                            }
                          ]
                        },
                        "response": {
                          "status": 200
                        },
                        "chaincode_id": {
                          "name": "mycc",
                          "version": "1.0"
                        }
                      }
                    },
                    "endorsements": [
                      {
                        "mspid": "Org1MSP",
                        "id_bytes": "3KT0HvuDK8DbXcfL195M8A=="
                      },
                      {
                        "mspid": "Org2MSP",
                        "id_bytes": "ypoN71ybKb2tE9FhCvJAcg=="
                      }
                    ]
                  }
                }
              }
            ]
          }
        }
      }
    ]
  },
  "metadata": {
    "metadata_config": {
      "value": {
        "index": 2
      }
    },
    "metadata_signature": {
      "signatures": [
        {
          "signature_header": {
            "mspid": "OrdererMSP",
            "id_bytes": "NRg5N7zJn/foUme9uALXeQ=="
          },
          "signature": "MEUCIQDKSfNfhS3XIgkgkX+mj9kmwcjICjpOoYfSB8rzXZWLnAIgG+NPXd9CTNcrdPcEo9GMFx9r+Ui+Kgv0veH3AWitWhc="
        }
      ]
    },
    "metadata_transactions_filter": "0 ",
    "metadata_orderer": {
      "value_raft": {
        "consenter_ids": [
          10,
          5,
          1,
          2,
          3,
          4,
          5,
          16,
          6,
          24,
          40
        ]
      }
    }
  }
}

```

## kown problem or limits

- the vaule of id_bytes is cut off.
- no test on idmex
- only to show HeaderType_ENDORSER_TRANSACTION and HeaderType_CONFIG
