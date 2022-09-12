# Synology NAS / Powerdown integration guide

This is guide is currently in draft state.

## Problem Statement

I would like my Synology NAS to shutdown before loadshedding cuts my power. I've heard of the EskomSePush API 
and I think I might be able to use it to make sure that my NAS is turned off 

## Solution 

Use the powerdown CLI **is-event-soon** command to access EskomSePush and determine whether loadshedding is imminent.

## Disclaimer

I have this working on a Synology NAS model [DS220+](https://www.synology.com/en-us/products/DS220+) and I am running 
[DSM 7.1-42661 Update 4](https://www.synology.com/en-global/releaseNote/DSM), 
I have only had an opportunity to test this on this configuration. Your mileage may vary.

## Prerequisites

### EskomSePush API Key

### Area Id

### Github Packages Personal Access Token

## Synology NAS Configuration

1. Install Docker
2. Enable SSH access
3. Log Docker into github packages
4. Disable SSH access
5. Enable a schedule task
6. 