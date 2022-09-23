# Resizable Cache

ResizableCache is a cache that can be resized. It is a wrapper around a cache that implements the Cache interface.

Its accomplishes this by keeping track of the number of items in the cache and the number of items that can be in the cache. 
A function determines to what size to resize the cache.


## Mechanism 

The resizable cache keeps two internal caches. but only when it is resizing while it fill in the second one.

### Resizing

Steps:
1. Create a new cache with the new size
2. Set pointer of `oldCache` to the same as `currentCache`
3. Set pointer of `currentCache` to the new cache
4. Traverse the old cache and add all items to the new cache, leaving out the expired items.
5. Set the pointer of `oldCache` to nil
