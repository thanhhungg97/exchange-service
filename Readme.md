    # Exchange Service with GoLang

# Feature

Exchange product between shope

Buy item from shope 
 
Import data to shope

# Idea
    The core has only two function:
        import and export
    Using event souring to store transaction 
    -> easy for update/ refresh current inventory if needed
    Other action like: by item , import data, exchange data base on that action
    
# Advantage
    Easy to switch with other database
    Using hexagon architechture
    -> Replace http controller with other techology like GRPC... 
    
# Limited
    This implemention inmemory only 
    but easy for switching with other database
    Only support 1 concurrent request at a time
    -> Need provide sharing by productid and locking mechainsm

    
# Future Improve
    Implementation with database
    Support concurrency.


# API Document
