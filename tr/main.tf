terraform{
    required_providers{
        mybotip = {
          //  source = "local/Users/nsysk_0101/univ/b4/tr/mybotip"
            version = "0.0.1" 
        }
    }
    required_version = "~> 1.8.1"
}

provider "mybotip"{
}