package main

type Config struct {
	DeviceName            string `json:"device_name"`
	UpdateFrequencySecond int    `json:"update_frequency_second"`
	RemoteServer          struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"remote_server"`
}
