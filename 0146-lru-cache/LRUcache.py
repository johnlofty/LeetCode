
class Node:
    def __init__(self, k, v):
        self.k = k
        self.v = v
        self.next = None
        self.prev = None

class LRUCache:

    def __init__(self, capacity: int):
        self.cap = capacity
        self.map = {}
        self.root = root = Node(None, None)
        root.prev = root.next = root

    def get(self, key: int) -> int:
        node = self.map.get(key)
        if not node:
            return -1
        else:
            prev_node = node.prev
            next_node = node.next
            next_node.prev = prev_node
            prev_node.next = next_node

            node.next = self.root
            node.prev = tail_node = self.root.prev
            self.root.prev = tail_node.next = node
            return node.v
        

    def put(self, key: int, value: int) -> None:
        node = self.map.get(key)
        if not node:
            node = Node(key, value)
            node.next = self.root
            node.prev = tail_node = self.root.prev
            self.root.prev = tail_node.next = node
            self.map[key] = node
            if len(self.map) > self.cap:
                self.pop()
        else:
            prev_node = node.prev
            next_node = node.next
            next_node.prev = prev_node
            prev_node.next = next_node

            node.next = self.root
            node.prev = tail_node = self.root.prev
            self.root.prev = tail_node.next = node
            node.v = value

    def pop(self):
        first_node = self.root.next
        next_node = first_node.next
        next_node.prev = self.root
        self.root.next = next_node
        self.map.pop(first_node.k)


