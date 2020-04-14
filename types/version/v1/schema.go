package v1

var OpenAPI = []byte(`
swagger: "2.0"
info:
  title: phenix
  version: "1.0"
paths:
  /topologies:
    get:
      responses:
        "200":
          description: array of topologies
          schema:
            type: array
            items:
              $ref: "#/definitions/Topology"
definitions:
  Topology:
    type: object
    required:
    - nodes
    properties:
      nodes:
        type: array
        items:
          $ref: "#/definitions/Node"
      vlans:
        type: array
        items:
          $ref: "#/definitions/VLAN"
  Node:
    type: object
    required:
    - type
    properties:
      type:
        type: string
        enum:
        - Firewall
        - Printer
        - Server
        - Switch
        - Router
        - SCEPTRE
        - VirtualMachine
  VLAN:
    type: object
    required:
    - name
    properties:
      name:
        type: string
`)

var schema = []byte(`
{
  "type": "object",
  "title": "SCEPTRE Topology",
  "required": [
    "nodes"
  ],
  "properties": {
    "nodes": {
      "$id": "#/properties/nodes",
      "type": "array",
      "title": "Nodes",
      "items": {
        "$id": "#/properties/nodes/items",
        "type": "object",
        "title": "Node",
        "headerTemplate": "{{i1}} - {{self.general.hostname}}",
        "required": [
          "type",
          "general",
          "hardware"
        ],
        "properties": {
          "type": {
            "$id": "#/properties/nodes/items/properties/type",
            "type": "string",
            "title": "Type",
            "enum": [
              "Firewall",
              "Printer",
              "Server",
              "Switch",
              "Router",
              "SCEPTRE",
              "VirtualMachine"
            ],
            "default": "VirtualMachine",
            "examples": [
              "Firewall",
              "Printer",
              "Server",
              "Switch",
              "Router",
              "SCEPTRE",
              "VirtualMachine"
            ],
            "pattern": "^(.*)$"
          },
          "general": {
            "$id": "#/properties/nodes/items/properties/general",
            "type": "object",
            "title": "General",
            "required": [
              "hostname"
            ],
            "properties": {
              "hostname": {
                "$id": "#/properties/nodes/items/properties/general/properties/hostname",
                "type": "string",
                "title": "Hostname",
                "minLength": 1,
                "default": "",
                "examples": [
                  "power-provider"
                ],
                "pattern": "^[\\w-]+$"
              },
              "description": {
                "$id": "#/properties/nodes/items/properties/general/properties/description",
                "type": "string",
                "title": "description",
                "default": "",
                "examples": [
                  "SCEPTRE power solver"
                ],
                "pattern": "^(.*)$"
              },
              "snapshot": {
                "$id": "#/properties/nodes/items/properties/general/properties/snapshot",
                "type": "boolean",
                "title": "snapshot",
                "default": false,
                "examples": [
                  false
                ]
              },
              "do_not_boot": {
                "$id": "#/properties/nodes/items/properties/general/properties/do_not_boot",
                "type": "boolean",
                "title": "do_not_boot",
                "default": false,
                "examples": [
                  false
                ]
              }
            }
          },
          "hardware": {
            "$id": "#/properties/nodes/items/properties/hardware",
            "type": "object",
            "title": "Hardware",
            "required": [
              "os_type",
              "drives"
            ],
            "properties": {
              "cpu": {
                "$id": "#/properties/nodes/items/properties/hardware/properties/cpu",
                "type": "string",
                "title": "cpu",
                "default": "Broadwell",
                "examples": [
                  "Broadwell",
                  "Haswell",
                  "core2duo",
                  "pentium3"
                ],
                "pattern": "^(.*)$"
              },
              "vcpus": {
                "$id": "#/properties/nodes/items/properties/hardware/properties/vcpus",
                "type": "string",
                "title": "vcpus",
                "default": "1",
                "examples": [
                  "4"
                ],
                "pattern": "^(.*)$"
              },
              "memory": {
                "$id": "#/properties/nodes/items/properties/hardware/properties/memory",
                "type": "string",
                "title": "memory",
                "default": "1024",
                "examples": [
                  "8192"
                ],
                "pattern": "^(.*)$"
              },
              "os_type": {
                "$id": "#/properties/nodes/items/properties/hardware/properties/os_type",
                "type": "string",
                "title": "os_type",
                "enum": [
                  "windows",
                  "linux",
                  "rhel",
                  "centos"
                ],
                "default": "linux",
                "examples": [
                  "windows",
                  "linux",
                  "rhel",
                  "centos"
                ],
                "pattern": "^(.*)$"
              },
              "drives": {
                "$id": "#/properties/nodes/items/properties/hardware/properties/drives",
                "type": "array",
                "title": "Drives",
                "items": {
                  "$id": "#/properties/nodes/items/properties/hardware/properties/drives/items",
                  "type": "object",
                  "title": "Drive",
                  "required": [
                    "image"
                  ],
                  "properties": {
                    "image": {
                      "$id": "#/properties/nodes/items/properties/hardware/properties/drives/items/properties/image",
                      "type": "string",
                      "title": "Image",
                      "minLength": 1,
                      "default": "",
                      "examples": [
                        "win10provider.qc2"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "interface": {
                      "$id": "#/properties/nodes/items/properties/hardware/properties/drives/items/properties/interface",
                      "type": "string",
                      "title": "interface",
                      "enum": [
                        "ahci",
                        "ide",
                        "scsi",
                        "sd",
                        "mtd",
                        "floppy",
                        "pflash",
                        "virtio"
                      ],
                      "default": "ide",
                      "examples": [
                        "ide"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "cache_mode": {
                      "$id": "#/properties/nodes/items/properties/hardware/properties/drives/items/properties/cache_mode",
                      "type": "string",
                      "title": "cache_mode",
                      "enum": [
                        "none",
                        "writeback",
                        "unsafe",
                        "directsync",
                        "writethrough"
                      ],
                      "default": "writeback",
                      "examples": [
                        "writeback"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "inject_partition": {
                      "$id": "#/properties/nodes/items/properties/hardware/properties/drives/items/properties/inject_partition",
                      "type": "string",
                      "title": "inject_partition",
                      "default": "1",
                      "examples": [
                        "2"
                      ],
                      "pattern": "^(.*)$"
                    }
                  }
                }
              }
            }
          },
          "network": {
            "$id": "#/properties/nodes/items/properties/network",
            "type": "object",
            "title": "Network",
            "required": [
              "interfaces"
            ],
            "properties": {
              "interfaces": {
                "$id": "#/properties/nodes/items/properties/network/properties/interfaces",
                "type": "array",
                "title": "Interfaces",
                "items": {
                  "$id": "#/properties/nodes/items/properties/network/properties/interfaces/items",
                  "type": "object",
                  "title": "Interface",
                  "oneOf": [
                    {
                      "$ref": "#/definitions/static_iface",
                      "title": "Static"
                    },
                    {
                      "$ref": "#/definitions/dhcp_iface",
                      "title": "DHCP"
                    },
                    {
                      "$ref": "#/definitions/serial_iface",
                      "title": "Serial"
                    }
                  ]
                }
              },
              "routes": {
                "$id": "#/properties/nodes/items/properties/network/properties/routes",
                "type": "array",
                "items": {
                  "$id": "#/properties/nodes/items/properties/network/properties/routes/items",
                  "type": "object",
                  "title": "Route",
                  "required": [
                    "destination",
                    "next",
                    "cost"
                  ],
                  "properties": {
                    "destination": {
                      "$id": "#/properties/nodes/items/properties/network/properties/routes/items/properties/destination",
                      "type": "string",
                      "title": "Destination",
                      "minLength": 1,
                      "default": "",
                      "examples": [
                        "192.168.0.0/24"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "next": {
                      "$id": "#/properties/nodes/items/properties/network/properties/routes/items/properties/next",
                      "type": "string",
                      "title": "Next",
                      "minLength": 1,
                      "default": "",
                      "examples": [
                        "192.168.1.254"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "cost": {
                      "$id": "#/properties/nodes/items/properties/network/properties/routes/items/properties/cost",
                      "type": "string",
                      "title": "Cost",
                      "minLength": 1,
                      "default": "1",
                      "examples": [
                        "1"
                      ],
                      "pattern": "^(.*)$"
                    }
                  }
                }
              },
              "ospf": {
                "$id": "#/properties/nodes/items/properties/network/properties/ospf",
                "type": "object",
                "title": "Ospf",
                "required": [
                  "router_id",
                  "areas"
                ],
                "properties": {
                  "router_id": {
                    "$id": "#/properties/nodes/items/properties/network/properties/ospf/properties/router_id",
                    "type": "string",
                    "title": "Router_id",
                    "default": "",
                    "examples": [
                      "0.0.0.1"
                    ],
                    "pattern": "^(.*)$"
                  },
                  "areas": {
                    "$id": "#/properties/nodes/items/properties/network/properties/ospf/properties/areas",
                    "type": "array",
                    "title": "Areas",
                    "items": {
                      "$id": "#/properties/nodes/items/properties/network/properties/ospf/properties/areas/items",
                      "type": "object",
                      "title": "Area",
                      "required": [
                        "area_id",
                        "area_networks"
                      ],
                      "properties": {
                        "area_id": {
                          "$id": "#/properties/nodes/items/properties/network/properties/ospf/properties/areas/items/properties/area_id",
                          "type": "string",
                          "title": "Area_id",
                          "default": "",
                          "examples": [
                            "0"
                          ],
                          "pattern": "^(.*)$"
                        },
                        "area_networks": {
                          "$id": "#/properties/nodes/items/properties/network/properties/ospf/properties/areas/items/properties/area_networks",
                          "type": "array",
                          "title": "Area_networks",
                          "items": {
                            "$id": "#/properties/nodes/items/properties/network/properties/ospf/properties/areas/items/properties/area_networks/items",
                            "type": "object",
                            "title": "Area Network",
                            "required": [
                              "network"
                            ],
                            "properties": {
                              "network": {
                                "$id": "#/properties/nodes/items/properties/network/properties/ospf/properties/areas/items/properties/area_networks/items/properties/network",
                                "type": "string",
                                "title": "Network",
                                "default": "",
                                "examples": [
                                  "10.1.25.0/24"
                                ],
                                "pattern": "^(.*)$"
                              }
                            }
                          }
                        }
                      }
                    }
                  }
                }
              },
              "rulesets": {
                "$id": "#/properties/nodes/items/properties/network/properties/rulesets",
                "type": "array",
                "title": "Rulesets",
                "items": {
                  "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items",
                  "type": "object",
                  "title": "Ruleset",
                  "required": [
                    "name",
                    "default",
                    "rules"
                  ],
                  "properties": {
                    "name": {
                      "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/name",
                      "type": "string",
                      "title": "Name",
                      "minLength": 1,
                      "default": "",
                      "examples": [
                        "OutToDMZ"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "description": {
                      "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/description",
                      "type": "string",
                      "title": "Description",
                      "default": "",
                      "examples": [
                        "From ICS to the DMZ network"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "default": {
                      "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/default",
                      "type": "string",
                      "title": "Default",
                      "minLength": 1,
                      "default": "",
                      "examples": [
                        "drop"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "rules": {
                      "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/rules",
                      "type": "array",
                      "title": "Rules",
                      "items": {
                        "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/rules/items",
                        "type": "object",
                        "title": "Rule",
                        "required": [
                          "id",
                          "action",
                          "protocol"
                        ],
                        "properties": {
                          "id": {
                            "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/rules/items/properties/id",
                            "type": "string",
                            "title": "Id",
                            "minLength": 1,
                            "default": "",
                            "examples": [
                              "10"
                            ],
                            "pattern": "^(.*)$"
                          },
                          "description": {
                            "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/rules/items/properties/description",
                            "type": "string",
                            "title": "Description",
                            "default": "",
                            "examples": [
                              "Allow UDP 10.1.26.80 ==> 10.2.25.0/24:123"
                            ],
                            "pattern": "^(.*)$"
                          },
                          "action": {
                            "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/rules/items/properties/action",
                            "type": "string",
                            "title": "Action",
                            "enum": [
                              "accept",
                              "drop",
                              "reject"
                            ],
                            "default": "drop",
                            "examples": [
                              "accept",
                              "drop",
                              "reject"
                            ],
                            "pattern": "^(.*)$"
                          },
                          "protocol": {
                            "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/rules/items/properties/protocol",
                            "type": "string",
                            "title": "Protocol",
                            "enum": [
                              "tcp",
                              "udp",
                              "icmp",
                              "esp",
                              "ah",
                              "all"
                            ],
                            "default": "tcp",
                            "examples": [
                              "tcp",
                              "udp",
                              "icmp",
                              "esp",
                              "ah",
                              "all"
                            ],
                            "pattern": "^(.*)$"
                          },
                          "source": {
                            "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/rules/items/properties/source",
                            "type": "object",
                            "title": "Source",
                            "required": [
                              "address"
                            ],
                            "properties": {
                              "address": {
                                "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/rules/items/properties/source/properties/address",
                                "type": "string",
                                "title": "Address",
                                "default": "",
                                "examples": [
                                  "10.1.24.0/24",
                                  "10.1.24.60"
                                ],
                                "pattern": "^(.*)$"
                              },
                              "port": {
                                "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/rules/items/properties/source/properties/port",
                                "type": "string",
                                "title": "Port",
                                "default": "",
                                "examples": [
                                  "3389"
                                ],
                                "pattern": "^(.*)$"
                              }
                            }
                          },
                          "destination": {
                            "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/rules/items/properties/destination",
                            "type": "object",
                            "title": "Destination",
                            "required": [
                              "address"
                            ],
                            "properties": {
                              "address": {
                                "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/rules/items/properties/source/properties/address",
                                "type": "string",
                                "title": "Address",
                                "default": "",
                                "examples": [
                                  "10.1.24.0/24",
                                  "10.1.24.60"
                                ],
                                "pattern": "^(.*)$"
                              },
                              "port": {
                                "$id": "#/properties/nodes/items/properties/network/properties/rulesets/items/properties/rules/items/properties/destination/properties/port",
                                "type": "string",
                                "title": "Port",
                                "default": "",
                                "examples": [
                                  "3389"
                                ],
                                "pattern": "^(.*)$"
                              }
                            }
                          }
                        }
                      }
                    }
                  }
                }
              }
            }
          },
          "injections": {
            "$id": "#/properties/nodes/items/properties/injections",
            "type": "array",
            "title": "Injections",
            "items": {
              "$id": "#/properties/nodes/items/properties/injections/items",
              "type": "object",
              "title": "Injection",
              "required": [
                "src",
                "dst"
              ],
              "properties": {
                "src": {
                  "$id": "#/properties/nodes/items/properties/injections/properties/src",
                  "type": "string",
                  "title": "Src",
                  "minLength": 1,
                  "default": "",
                  "examples": [
                    "ACTIVSg2000.PWB"
                  ],
                  "pattern": "^(.*)$"
                },
                "dst": {
                  "$id": "#/properties/nodes/items/properties/injections/properties/dst",
                  "type": "string",
                  "title": "Dst",
                  "minLength": 1,
                  "default": "",
                  "examples": [
                    "sceptre/ACTIVSg2000.PWB"
                  ],
                  "pattern": "^(.*)$"
                },
                "description": {
                  "$id": "#/properties/nodes/items/properties/injections/properties/description",
                  "type": "string",
                  "title": "description",
                  "default": "",
                  "examples": [
                    "PowerWorld case binary data"
                  ],
                  "pattern": "^(.*)$"
                }
              }
            }
          },
          "metadata": {
            "$id": "#/properties/nodes/items/properties/metadata",
            "type": "object",
            "title": "Metadata",
            "properties": {
              "infrastructure": {
                "$id": "#/properties/nodes/items/properties/metadata/properties/infrastructure",
                "type": "string",
                "title": "Infrastructure",
                "enum": [
                  "power-transmission",
                  "power-distribution",
                  "batch-process"
                ],
                "default": "power-transmission",
                "examples": [
                  "power-transmission",
                  "power-distribution",
                  "batch-process"
                ],
                "pattern": "^(.*)$"
              },
              "provider": {
                "$id": "#/properties/nodes/items/properties/metadata/properties/provider",
                "type": "string",
                "title": "Provider",
                "default": "power-provider",
                "examples": [
                  "power-provider",
                  "simulink-provider"
                ],
                "pattern": "^(.*)$"
              },
              "simulator": {
                "$id": "#/properties/nodes/items/properties/metadata/properties/simulator",
                "type": "string",
                "title": "Simulator",
                "enum": [
                  "Dummy",
                  "PSSE",
                  "PyPower",
                  "PowerWorld",
                  "PowerWorldDynamics",
                  "OpenDSS",
                  "Simulink"
                ],
                "default": "PowerWorld",
                "examples": [
                  "PowerWorld"
                ],
                "pattern": "^(.*)$"
              },
              "publish_endpoint": {
                "$id": "#/properties/nodes/items/properties/metadata/properties/publish_endpoint",
                "type": "string",
                "title": "Publish_endpoint",
                "default": "udp://*;239.0.0.1:40000",
                "examples": [
                  "udp://*;239.0.0.1:40000"
                ],
                "pattern": "^(.*)$"
              },
              "cycle_time": {
                "$id": "#/properties/nodes/items/properties/metadata/properties/cycle_time",
                "type": "string",
                "title": "Cycle_time",
                "default": "500",
                "examples": [
                  "1000"
                ],
                "pattern": "^(.*)$"
              },
              "dnp3": {
                "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3",
                "type": "array",
                "title": "Dnp3",
                "items": {
                  "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items",
                  "type": "object",
                  "title": "DNP3 Metadata",
                  "required": [
                    "type",
                    "name"
                  ],
                  "properties": {
                    "type": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/type",
                      "type": "string",
                      "title": "Type",
                      "default": "",
                      "examples": [
                        "bus"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "name": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/name",
                      "type": "string",
                      "title": "Name",
                      "default": "",
                      "examples": [
                        "bus-2052"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "analog-read": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/analog-read",
                      "type": "array",
                      "title": "The Analog-read Schema",
                      "items": {
                        "$id": "#/properties/metadata/properties/dnp3/items/properties/analog-read/items",
                        "type": "object",
                        "title": "The Items Schema",
                        "required": [
                          "field",
                          "register_number",
                          "register_type"
                        ],
                        "properties": {
                          "field": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/analog-read/items/properties/field",
                            "type": "string",
                            "title": "The Field Schema",
                            "default": "",
                            "examples": [
                              "voltage"
                            ],
                            "pattern": "^(.*)$"
                          },
                          "register_number": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/analog-read/items/properties/register_number",
                            "type": "integer",
                            "title": "The Register_number Schema",
                            "default": 0,
                            "examples": [
                              0
                            ]
                          },
                          "register_type": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/analog-read/items/properties/register_type",
                            "type": "string",
                            "title": "The Register_type Schema",
                            "default": "",
                            "examples": [
                              "analog-input"
                            ],
                            "pattern": "^(.*)$"
                          }
                        }
                      }
                    },
                    "binary-read": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/binary-read",
                      "type": "array",
                      "title": "The Binary-read Schema",
                      "items": {
                        "$id": "#/properties/metadata/properties/dnp3/items/properties/binary-read/items",
                        "type": "object",
                        "title": "The Items Schema",
                        "required": [
                          "field",
                          "register_number",
                          "register_type"
                        ],
                        "properties": {
                          "field": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/binary-read/items/properties/field",
                            "type": "string",
                            "title": "The Field Schema",
                            "default": "",
                            "examples": [
                              "active"
                            ],
                            "pattern": "^(.*)$"
                          },
                          "register_number": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/binary-read/items/properties/register_number",
                            "type": "integer",
                            "title": "The Register_number Schema",
                            "default": 0,
                            "examples": [
                              7
                            ]
                          },
                          "register_type": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/binary-read/items/properties/register_type",
                            "type": "string",
                            "title": "The Register_type Schema",
                            "default": "",
                            "examples": [
                              "binary-input"
                            ],
                            "pattern": "^(.*)$"
                          }
                        }
                      }
                    },
                    "binary-read-write": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/binary-read-write",
                      "type": "array",
                      "title": "The Binary-read-write Schema",
                      "items": {
                        "$id": "#/properties/metadata/properties/dnp3/items/properties/binary-read-write/items",
                        "type": "object",
                        "title": "The Items Schema",
                        "required": [
                          "field",
                          "register_number",
                          "register_type"
                        ],
                        "properties": {
                          "field": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/binary-read-write/items/properties/field",
                            "type": "string",
                            "title": "The Field Schema",
                            "default": "",
                            "examples": [
                              "active"
                            ],
                            "pattern": "^(.*)$"
                          },
                          "register_number": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/binary-read-write/items/properties/register_number",
                            "type": "integer",
                            "title": "The Register_number Schema",
                            "default": 0,
                            "examples": [
                              9
                            ]
                          },
                          "register_type": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3/items/properties/binary-read-write/items/properties/register_type",
                            "type": "string",
                            "title": "The Register_type Schema",
                            "default": "",
                            "examples": [
                              "binary-output"
                            ],
                            "pattern": "^(.*)$"
                          }
                        }
                      }
                    }
                  }
                }
              },
              "dnp3-serial": {
                "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial",
                "type": "array",
                "title": "Dnp3-serial",
                "items": {
                  "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items",
                  "type": "object",
                  "title": "DNP3-serial Metadata",
                  "required": [
                    "type",
                    "name"
                  ],
                  "properties": {
                    "type": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/type",
                      "type": "string",
                      "title": "Type",
                      "default": "",
                      "examples": [
                        "bus"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "name": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/name",
                      "type": "string",
                      "title": "Name",
                      "default": "",
                      "examples": [
                        "bus-2052"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "analog-read": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/analog-read",
                      "type": "array",
                      "title": "The Analog-read Schema",
                      "items": {
                        "$id": "#/properties/metadata/properties/dnp3-serial/items/properties/analog-read/items",
                        "type": "object",
                        "title": "The Items Schema",
                        "required": [
                          "field",
                          "register_number",
                          "register_type"
                        ],
                        "properties": {
                          "field": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/analog-read/items/properties/field",
                            "type": "string",
                            "title": "The Field Schema",
                            "default": "",
                            "examples": [
                              "voltage"
                            ],
                            "pattern": "^(.*)$"
                          },
                          "register_number": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/analog-read/items/properties/register_number",
                            "type": "integer",
                            "title": "The Register_number Schema",
                            "default": 0,
                            "examples": [
                              0
                            ]
                          },
                          "register_type": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/analog-read/items/properties/register_type",
                            "type": "string",
                            "title": "The Register_type Schema",
                            "default": "",
                            "examples": [
                              "analog-input"
                            ],
                            "pattern": "^(.*)$"
                          }
                        }
                      }
                    },
                    "binary-read": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/binary-read",
                      "type": "array",
                      "title": "The Binary-read Schema",
                      "items": {
                        "$id": "#/properties/metadata/properties/dnp3-serial/items/properties/binary-read/items",
                        "type": "object",
                        "title": "The Items Schema",
                        "required": [
                          "field",
                          "register_number",
                          "register_type"
                        ],
                        "properties": {
                          "field": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/binary-read/items/properties/field",
                            "type": "string",
                            "title": "The Field Schema",
                            "default": "",
                            "examples": [
                              "active"
                            ],
                            "pattern": "^(.*)$"
                          },
                          "register_number": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/binary-read/items/properties/register_number",
                            "type": "integer",
                            "title": "The Register_number Schema",
                            "default": 0,
                            "examples": [
                              7
                            ]
                          },
                          "register_type": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/binary-read/items/properties/register_type",
                            "type": "string",
                            "title": "The Register_type Schema",
                            "default": "",
                            "examples": [
                              "binary-input"
                            ],
                            "pattern": "^(.*)$"
                          }
                        }
                      }
                    },
                    "binary-read-write": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/binary-read-write",
                      "type": "array",
                      "title": "The Binary-read-write Schema",
                      "items": {
                        "$id": "#/properties/metadata/properties/dnp3-serial/items/properties/binary-read-write/items",
                        "type": "object",
                        "title": "The Items Schema",
                        "required": [
                          "field",
                          "register_number",
                          "register_type"
                        ],
                        "properties": {
                          "field": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/binary-read-write/items/properties/field",
                            "type": "string",
                            "title": "The Field Schema",
                            "default": "",
                            "examples": [
                              "active"
                            ],
                            "pattern": "^(.*)$"
                          },
                          "register_number": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/binary-read-write/items/properties/register_number",
                            "type": "integer",
                            "title": "The Register_number Schema",
                            "default": 0,
                            "examples": [
                              9
                            ]
                          },
                          "register_type": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/dnp3-serial/items/properties/binary-read-write/items/properties/register_type",
                            "type": "string",
                            "title": "The Register_type Schema",
                            "default": "",
                            "examples": [
                              "binary-output"
                            ],
                            "pattern": "^(.*)$"
                          }
                        }
                      }
                    }
                  }
                }
              },
              "modbus": {
                "$id": "#/properties/nodes/items/properties/metadata/properties/modbus",
                "type": "array",
                "title": "Modbus",
                "items": {
                  "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items",
                  "type": "object",
                  "title": "Modbus Metadata",
                  "required": [
                    "type",
                    "name"
                  ],
                  "properties": {
                    "type": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/type",
                      "type": "string",
                      "title": "Type",
                      "default": "",
                      "examples": [
                        "bus"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "name": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/name",
                      "type": "string",
                      "title": "Name",
                      "default": "",
                      "examples": [
                        "bus-2052"
                      ],
                      "pattern": "^(.*)$"
                    },
                    "analog-read": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/analog-read",
                      "type": "array",
                      "title": "The Analog-read Schema",
                      "items": {
                        "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/analog-read/items",
                        "type": "object",
                        "title": "The Items Schema",
                        "required": [
                          "field",
                          "register_number",
                          "register_type"
                        ],
                        "properties": {
                          "field": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/analog-read/items/properties/field",
                            "type": "string",
                            "title": "The Field Schema",
                            "default": "",
                            "examples": [
                              "voltage"
                            ],
                            "pattern": "^(.*)$"
                          },
                          "register_number": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/analog-read/items/properties/register_number",
                            "type": "integer",
                            "title": "The Register_number Schema",
                            "default": 0,
                            "examples": [
                              30000
                            ]
                          },
                          "register_type": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/analog-read/items/properties/register_type",
                            "type": "string",
                            "title": "The Register_type Schema",
                            "default": "",
                            "examples": [
                              "input-register"
                            ],
                            "pattern": "^(.*)$"
                          }
                        }
                      }
                    },
                    "binary-read": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/binary-read",
                      "type": "array",
                      "title": "The Binary-read Schema",
                      "items": {
                        "$id": "#/properties/metadata/properties/modbus/items/properties/binary-read/items",
                        "type": "object",
                        "title": "The Items Schema",
                        "required": [
                          "field",
                          "register_number",
                          "register_type"
                        ],
                        "properties": {
                          "field": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/binary-read/items/properties/field",
                            "type": "string",
                            "title": "The Field Schema",
                            "default": "",
                            "examples": [
                              "active"
                            ],
                            "pattern": "^(.*)$"
                          },
                          "register_number": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/binary-read/items/properties/register_number",
                            "type": "integer",
                            "title": "The Register_number Schema",
                            "default": 0,
                            "examples": [
                              10000
                            ]
                          },
                          "register_type": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/binary-read/items/properties/register_type",
                            "type": "string",
                            "title": "The Register_type Schema",
                            "default": "",
                            "examples": [
                              "discrete-input"
                            ],
                            "pattern": "^(.*)$"
                          }
                        }
                      }
                    },
                    "binary-read-write": {
                      "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/binary-read-write",
                      "type": "array",
                      "title": "The Binary-read-write Schema",
                      "items": {
                        "$id": "#/properties/metadata/properties/modbus/items/properties/binary-read-write/items",
                        "type": "object",
                        "title": "The Items Schema",
                        "required": [
                          "field",
                          "register_number",
                          "register_type"
                        ],
                        "properties": {
                          "field": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/binary-read-write/items/properties/field",
                            "type": "string",
                            "title": "The Field Schema",
                            "default": "",
                            "examples": [
                              "active"
                            ],
                            "pattern": "^(.*)$"
                          },
                          "register_number": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/binary-read-write/items/properties/register_number",
                            "type": "integer",
                            "title": "The Register_number Schema",
                            "default": 0,
                            "examples": [
                              0
                            ]
                          },
                          "register_type": {
                            "$id": "#/properties/nodes/items/properties/metadata/properties/modbus/items/properties/binary-read-write/items/properties/register_type",
                            "type": "string",
                            "title": "The Register_type Schema",
                            "default": "",
                            "examples": [
                              "coil"
                            ],
                            "pattern": "^(.*)$"
                          }
                        }
                      }
                    }
                  }
                }
              },
              "logic": {
                "$id": "#/properties/nodes/items/properties/metadata/properties/logic",
                "type": "string",
                "title": "Logic",
                "default": "",
                "examples": [
                  "Tank1.fill_control = Tank1.tank_volume < Tank1.level_setpoint || (Tank1.tank_volume < 1.5*Tank1.level_setpoint && Tank1.fill_control == 1); Pump1.control = ! FillingStation1.request == 0 && Tank1.tank_volume>0; Pump1.active = 1==1"
                ],
                "pattern": "^(.*)$"
              },
              "connected_rtus": {
                "$id": "#/properties/nodes/items/properties/metadata/properties/connected_rtus",
                "type": "array",
                "title": "Connected RTUs",
                "items": {
                  "$id": "#/properties/nodes/items/properties/metadata/properties/connected_rtus/items",
                  "type": "string",
                  "title": "RTU",
                  "default": "",
                  "examples": [
                    "rtu-1",
                    "fuel-rtu-2"
                  ],
                  "pattern": "^(.*)$"
                }
              },
              "connect_to_scada": {
                "type": "boolean",
                "title": "connect_to_scada",
                "default": false,
                "examples": [
                  true,
                  false
                ]
              },
              "manual_register_config": {
                "$id": "#/properties/nodes/items/properties/metadata/properties/manual_register_config",
                "type": "string",
                "title": "Manual register configuration",
                "default": "false",
                "examples": [
                  "false"
                ]
              }
            }
          }
        }
      }
    },
    "vlans": {
      "$id": "#/properties/vlans",
      "type": "array",
      "title": "VLANs",
      "items": {
        "$id": "#/properties/vlans/items",
        "type": "object",
        "title": "VLAN",
        "required": [
          "name",
          "id"
        ],
        "properties": {
          "name": {
            "$id": "#/properties/vlans/items/properties/name",
            "type": "string",
            "title": "Name",
            "minLength": 1,
            "default": "",
            "examples": [
              "SCADAPWR"
            ],
            "pattern": "^(.*)$"
          },
          "id": {
            "$id": "#/properties/vlans/items/properties/id",
            "type": "string",
            "title": "Id",
            "minLength": 1,
            "default": "",
            "examples": [
              "101"
            ],
            "pattern": "^(.*)$"
          }
        }
      }
    }
  },
  "definitions": {
    "iface": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "Name",
          "default": "",
          "examples": [
            "eth0"
          ],
          "pattern": "^[\\w-]+$"
        },
        "vlan": {
          "type": "string",
          "title": "VLAN",
          "default": "",
          "examples": [
            "SCADAPWR"
          ],
          "pattern": "^[\\w-]+$"
        },
        "autostart": {
          "type": "boolean",
          "title": "autostart",
          "default": true,
          "examples": [
            true,
            false
          ]
        },
        "mac": {
          "type": "string",
          "title": "mac",
          "default": "",
          "examples": [
            "00:11:22:33:44:55:66"
          ],
          "pattern": "^([0-9a-fA-F]{2}[:-]){5}([0-9a-fA-F]){2}$"
        },
        "mtu": {
          "type": "integer",
          "title": "mtu",
          "default": 1500,
          "examples": [
            1500
          ]
        }
      },
      "required": [
        "name",
        "vlan"
      ]
    },
    "iface_address": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string",
          "title": "address",
          "default": "",
          "examples": [
            "192.168.1.100"
          ],
          "pattern": "^(?!0)(?!.*\\.$)((1?\\d?\\d|25[0-5]|2[0-4]\\d)(\\.|$)){4}$"
        },
        "mask": {
          "type": "integer",
          "title": "mask",
          "default": 24,
          "minimum": 0,
          "maximum": 32,
          "examples": [
            24
          ]
        },
        "gateway": {
          "type": "string",
          "title": "gateway",
          "default": "",
          "examples": [
            "192.168.1.1"
          ],
          "pattern": "^(?!0)(?!.*\\.$)((1?\\d?\\d|25[0-5]|2[0-4]\\d)(\\.|$)){4}$"
        }
      },
      "required": [
        "address",
        "mask"
      ]
    },
    "iface_rulesets": {
      "type": "object",
      "properties": {
        "ruleset_out": {
          "type": "string",
          "title": "ruleset_out",
          "default": "",
          "examples": [
            "OutToICS"
          ],
          "pattern": "^[\\w-]+$"
        },
        "ruleset_in": {
          "type": "string",
          "title": "ruleset_in",
          "default": "",
          "examples": [
            "InFromSCADA"
          ],
          "pattern": "^[\\w-]+$"
        }
      }
    },
    "static_iface": {
      "allOf": [
        {
          "$ref": "#/definitions/iface"
        },
        {
          "$ref": "#/definitions/iface_address"
        },
        {
          "$ref": "#/definitions/iface_rulesets"
        },
        {
          "properties": {
            "type": {
              "type": "string",
              "title": "type",
              "enum": [
                "ethernet"
              ],
              "default": "ethernet",
              "examples": [
                "ethernet"
              ]
            },
            "proto": {
              "type": "string",
              "title": "proto",
              "enum": [
                "static",
                "ospf"
              ],
              "default": "static",
              "examples": [
                "static",
                "ospf"
              ]
            }
          },
          "required": [
            "type",
            "proto"
          ]
        }
      ]
    },
    "dhcp_iface": {
      "allOf": [
        {
          "$ref": "#/definitions/iface"
        },
        {
          "$ref": "#/definitions/iface_rulesets"
        },
        {
          "properties": {
            "type": {
              "type": "string",
              "title": "type",
              "enum": [
                "ethernet"
              ],
              "default": "ethernet",
              "examples": [
                "ethernet"
              ]
            },
            "proto": {
              "type": "string",
              "title": "proto",
              "enum": [
                "dhcp"
              ],
              "default": "dhcp",
              "examples": [
                "dhcp"
              ]
            }
          },
          "required": [
            "type",
            "proto"
          ]
        }
      ]
    },
    "serial_iface": {
      "allOf": [
        {
          "$ref": "#/definitions/iface"
        },
        {
          "$ref": "#/definitions/iface_address"
        },
        {
          "$ref": "#/definitions/iface_rulesets"
        },
        {
          "properties": {
            "type": {
              "type": "string",
              "title": "type",
              "enum": [
                "serial"
              ],
              "default": "serial",
              "examples": [
                "serial"
              ]
            },
            "proto": {
              "type": "string",
              "title": "proto",
              "enum": [
                "static"
              ],
              "default": "static",
              "examples": [
                "static"
              ]
            },
            "udp_port": {
              "type": "integer",
              "title": "udp_port",
              "default": 8989,
              "minimum": 0,
              "maximum": 65535,
              "examples": [
                8989
              ]
            },
            "baud_rate": {
              "type": "integer",
              "title": "baud_rate",
              "enum": [
                110,
                300,
                600,
                1200,
                2400,
                4800,
                9600,
                14400,
                19200,
                38400,
                57600,
                115200,
                128000,
                256000
              ],
              "default": 9600,
              "examples": [
                9600
              ]
            },
            "device": {
              "type": "string",
              "title": "device",
              "default": "/dev/ttyS0",
              "examples": [
                "/dev/ttyS0"
              ],
              "pattern": "^[\\w\\/]+\\w+$"
            }
          },
          "required": [
            "type",
            "proto",
            "udp_port",
            "baud_rate",
            "device"
          ]
        }
      ]
    }
  }
}
`)
