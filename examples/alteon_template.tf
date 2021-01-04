resource "azurerm_template_deployment" "alteon" {
  name                = "alteon-deploy"
  resource_group_name = "alteon-group"
  deployment_mode     = "Incremental"
  template_body    = <<TEMPLATE
{
  "$schema": "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
  "contentVersion": "1.0.0.0",
  "parameters": {
        "location": {
            "value": "eastus"
        },
        "networkInterfaceName": {
            "value": "alteon-test20"
        },
        "networkSecurityGroupName": {
            "value": "alteon-test-nsg"
        },
        "networkSecurityGroupRules": {
            "value": [
                {
                    "name": "mgmt-https-singleIP",
                    "properties": {
                        "priority": 1010,
                        "protocol": "TCP",
                        "access": "Allow",
                        "direction": "Inbound",
                        "sourceApplicationSecurityGroups": [],
                        "destinationApplicationSecurityGroups": [],
                        "sourceAddressPrefix": "*",
                        "sourcePortRange": "*",
                        "destinationAddressPrefix": "*",
                        "destinationPortRange": "8443"
                    }
                },
                {
                    "name": "mgmt-https-multi-IP",
                    "properties": {
                        "priority": 1020,
                        "protocol": "TCP",
                        "access": "Allow",
                        "direction": "Inbound",
                        "sourceApplicationSecurityGroups": [],
                        "destinationApplicationSecurityGroups": [],
                        "sourceAddressPrefix": "*",
                        "sourcePortRange": "*",
                        "destinationAddressPrefix": "*",
                        "destinationPortRange": "443"
                    }
                },
                {
                    "name": "SingleIP-GSLB-DSSP",
                    "properties": {
                        "priority": 1030,
                        "protocol": "TCP",
                        "access": "Allow",
                        "direction": "Inbound",
                        "sourceApplicationSecurityGroups": [],
                        "destinationApplicationSecurityGroups": [],
                        "sourceAddressPrefix": "*",
                        "sourcePortRange": "*",
                        "destinationAddressPrefix": "*",
                        "destinationPortRange": "4480"
                    }
                },
                {
                    "name": "Local-License-Server",
                    "properties": {
                        "priority": 1040,
                        "protocol": "TCP",
                        "access": "Allow",
                        "direction": "Inbound",
                        "sourceApplicationSecurityGroups": [],
                        "destinationApplicationSecurityGroups": [],
                        "sourceAddressPrefix": "*",
                        "sourcePortRange": "*",
                        "destinationAddressPrefix": "*",
                        "destinationPortRange": "7070"
                    }
                },
                {
                    "name": "default-allow-ssh",
                    "properties": {
                        "priority": 1050,
                        "protocol": "TCP",
                        "access": "Allow",
                        "direction": "Inbound",
                        "sourceApplicationSecurityGroups": [],
                        "destinationApplicationSecurityGroups": [],
                        "sourceAddressPrefix": "*",
                        "sourcePortRange": "*",
                        "destinationAddressPrefix": "*",
                        "destinationPortRange": "22"
                    }
                }
            ]
        },
        "subnetName": {
            "value": "net-mgmt"
        },
        "virtualNetworkId": {
            "value": "/subscriptions/8b308087-ea81-4958-b3be-c9360a994c53/resourceGroups/net-engineering/providers/Microsoft.Network/virtualNetworks/VNET-East-Pa"
        },
        "publicIpAddressName": {
            "value": "alteon-test-ip"
        },
        "publicIpAddressType": {
            "value": "Dynamic"
        },
        "publicIpAddressSku": {
            "value": "Basic"
        },
        "virtualMachineName": {
            "value": "alteon-test"
        },
        "virtualMachineComputerName": {
            "value": "alteon-test"
        },
        "virtualMachineRG": {
            "value": "net-engineering"
        },
        "osDiskType": {
            "value": "StandardSSD_LRS"
        },
        "virtualMachineSize": {
            "value": "Standard_DS1_v2"
        },
        "adminUsername": {
            "value": "azureuser"
        },
        "adminPassword": {
            "value": null
        },
        "autoShutdownStatus": {
            "value": "Enabled"
        },
        "autoShutdownTime": {
            "value": "19:00"
        },
        "autoShutdownTimeZone": {
            "value": "UTC"
        },
        "autoShutdownNotificationStatus": {
            "value": "Enabled"
        },
        "autoShutdownNotificationLocale": {
            "value": "en"
        },
        "autoShutdownNotificationEmail": {
            "value": "iromaniuk@advisor360test.com"
        }
    },
  "variables": {
        "nsgId": "[resourceId(resourceGroup().name, 'Microsoft.Network/networkSecurityGroups', parameters('networkSecurityGroupName'))]",
        "vnetId": "[parameters('virtualNetworkId')]",
        "subnetRef": "[concat(variables('vnetId'), '/subnets/', parameters('subnetName'))]"
    },
  "resources": [
        {
            "type": "Microsoft.Network/networkInterfaces",
            "apiVersion": "2018-10-01",
            "name": "[parameters('networkInterfaceName')]",
            "location": "[parameters('location')]",
            "dependsOn": [
                "[concat('Microsoft.Network/networkSecurityGroups/', parameters('networkSecurityGroupName'))]",
                "[concat('Microsoft.Network/publicIpAddresses/', parameters('publicIpAddressName'))]"
            ],
            "properties": {
                "ipConfigurations": [
                    {
                        "name": "ipconfig1",
                        "properties": {
                            "subnet": {
                                "id": "[variables('subnetRef')]"
                            },
                            "privateIPAllocationMethod": "Dynamic",
                            "publicIpAddress": {
                                "id": "[resourceId(resourceGroup().name, 'Microsoft.Network/publicIpAddresses', parameters('publicIpAddressName'))]"
                            }
                        }
                    }
                ],
                "networkSecurityGroup": {
                    "id": "[variables('nsgId')]"
                }
            }
        },
        {
            "type": "Microsoft.Network/networkSecurityGroups",
            "apiVersion": "2019-02-01",
            "name": "[parameters('networkSecurityGroupName')]",
            "location": "[parameters('location')]",
            "properties": {
                "securityRules": "[parameters('networkSecurityGroupRules')]"
            }
        },
        {
            "type": "Microsoft.Network/publicIpAddresses",
            "apiVersion": "2019-02-01",
            "name": "[parameters('publicIpAddressName')]",
            "location": "[parameters('location')]",
            "sku": {
                "name": "[parameters('publicIpAddressSku')]"
            },
            "properties": {
                "publicIpAllocationMethod": "[parameters('publicIpAddressType')]"
            }
        },
        {
            "type": "Microsoft.Compute/virtualMachines",
            "apiVersion": "2020-06-01",
            "name": "[parameters('virtualMachineName')]",
            "location": "[parameters('location')]",
            "dependsOn": [
                "[concat('Microsoft.Network/networkInterfaces/', parameters('networkInterfaceName'))]"
            ],
            "plan": {
                "name": "radware-alteon-ng-va-adc",
                "publisher": "radware",
                "product": "radware-alteon-va"
            },
            "properties": {
                "hardwareProfile": {
                    "vmSize": "[parameters('virtualMachineSize')]"
                },
                "storageProfile": {
                    "osDisk": {
                        "createOption": "fromImage",
                        "managedDisk": {
                            "storageAccountType": "[parameters('osDiskType')]"
                        }
                    },
                    "imageReference": {
                        "publisher": "radware",
                        "offer": "radware-alteon-va",
                        "sku": "radware-alteon-ng-va-adc",
                        "version": "32.6.10000207"
                    }
                },
                "networkProfile": {
                    "networkInterfaces": [
                        {
                            "id": "[resourceId('Microsoft.Network/networkInterfaces', parameters('networkInterfaceName'))]"
                        }
                    ]
                },
                "osProfile": {
                    "computerName": "[parameters('virtualMachineComputerName')]",
                    "adminUsername": "[parameters('adminUsername')]",
                    "adminPassword": "[parameters('adminPassword')]"
                },
                "diagnosticsProfile": {
                    "bootDiagnostics": {
                        "enabled": true
                    }
                }
            }
        },
        {
            "type": "Microsoft.DevTestLab/schedules",
            "apiVersion": "2017-04-26-preview",
            "name": "[concat('shutdown-computevm-', parameters('virtualMachineName'))]",
            "location": "[parameters('location')]",
            "dependsOn": [
                "[concat('Microsoft.Compute/virtualMachines/', parameters('virtualMachineName'))]"
            ],
            "properties": {
                "status": "[parameters('autoShutdownStatus')]",
                "taskType": "ComputeVmShutdownTask",
                "dailyRecurrence": {
                    "time": "[parameters('autoShutdownTime')]"
                },
                "timeZoneId": "[parameters('autoShutdownTimeZone')]",
                "targetResourceId": "[resourceId('Microsoft.Compute/virtualMachines', parameters('virtualMachineName'))]",
                "notificationSettings": {
                    "status": "[parameters('autoShutdownNotificationStatus')]",
                    "notificationLocale": "[parameters('autoShutdownNotificationLocale')]",
                    "timeInMinutes": "30",
                    "emailRecipient": "[parameters('autoShutdownNotificationEmail')]"
                }
            }
        }
    ]
}
TEMPLATE

  // NOTE: whilst we show an inline template here, we recommend
  // sourcing this from a file for readability/editor support
}