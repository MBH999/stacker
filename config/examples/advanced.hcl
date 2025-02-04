#environments {
#  environment "connectivity" {
#    tags = []
#    deploy_as_stack = false
#  }
#  environment "identity" {
#    tags = []
#    deploy_as_stack = false
#  }
#  environment "management" {
#    tags = []
#    deploy_as_stack = false
#  }
#}

regions {
  region "uks" {
    tags = []
    deploy_as_stack = true
  }
  region "ukw" {
    tags = []
    deploy_as_stack = true
  }
}

stacks {
  stack "connectivity" {
    tags = []
    stack "virtual_wan" {
      tags = []
    }

    stack "virtual_hub" {
      tags = []
    }

    stack "palo_alto" {
      tags = []
    }

    stack "bastion" {
      tags = []
    }

    stack "dns_resolver" {
      tags = []
    }

    stack "virtual_hub_connections" {
      tags = []
    }

    stack "log_analytics" {
      tags = []
    }

    stack "virtual_networks" {
      tags = []
    }
  }
  
  stack "identity" {
    tags = []
    stack "virtual_networks" {
      tags = []
    }

    stack "dc_dev" {
      tags = []
      stack "application_security_group" {
        tags = []
      }
      stack "log_analytics" {
        tags = []
      }
      stack "recovery_services_vault" {
        tags = []
      }
      stack "network_security_group" {
        tags = []
      }
    }

    stack "dc_tst" {
      tags = []
      stack "application_security_group" {
        tags = []
      }
      stack "log_analytics" {
        tags = []
      }
      stack "recovery_services_vault" {
        tags = []
      }
      stack "network_security_group" {
        tags = []
      }
    }

    stack "dc_prd" {
      tags = []
      stack "application_security_group" {
        tags = []
      }
      stack "log_analytics" {
        tags = []
      }
      stack "recovery_services_vault" {
        tags = []
      }
      stack "network_security_group" {
        tags = []
      }
    }

    stack "kv_tst" {
      tags = []
    }
  }

  stack "management" {
    tags = []
    exclude_regions = ["ukw"]
    stack "virtual_networks" {
      tags = []
    }

    stack "build_agent" {
      tags = []
      exclude_environments = ["identity", "connectivity"]
    }

    stack "vm_access_kv" {
      tags = []
      exclude_environments = ["identity", "connectivity"]
    }
  }
}


