#!/usr/bin/env python
#
# Generate messages
#
# {
#   "id": "LBIT",
#   "t": "2021-06-17T22:13:37.166925",
#   "type": "start",
#   "data": {
#     "pos": 110777
#   }
# }


import argparse
import datetime
import json
import random
import string


def randstring(length=4):
    return "".join((random.choice(string.ascii_uppercase) for _ in range(length)))


def generate_messages(n):
    types = ["start", "stop", "mark"]

    for v in range(args.n):
        msg_type = types[v % len(types)]
        msg = {
            "id": randstring(),
            "t": datetime.datetime.now().strftime("%Y%m%d%H%m%S"),
            "type": msg_type,
        }
        if msg_type == "start":
            data = {
                "pos": random.randint(0, 999999),
            }
        elif msg_type == "stop":
            data = {
                "pos": random.randint(0, 999999),
                "reason": randstring(length=1),
            }
        elif msg_type == "mark":
            data = {
                "pos": random.randint(0, 999999),
                "color": random.randint(0, 16),
            }

        msg["data"] = data
        yield msg


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-n", type=int, default=10, help="messages to generate")

    args = parser.parse_args()

    for msg in generate_messages(args.n):
        print(json.dumps(msg))
