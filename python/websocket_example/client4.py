import asyncio

import websockets


async def handle_message(player_id, ws, message):
    if message.startswith("UpdateScene:"):
        # 处理更新场景的消息
        scene_data = message[len("UpdateScene:"):]
        print(f"Player {player_id} received scene update: {scene_data}")
    elif message.startswith("PlayerAction:"):
        # 处理玩家操作的消息
        action_data = message[len("PlayerAction:"):]
        print(f"Player {player_id} received player action: {action_data}")
        # 在这里处理玩家操作的逻辑


async def send_heartbeat(ws):
    while True:
        # 每 2 秒发送一次心跳消息
        await asyncio.sleep(2)
        print("Sending heartbeat...")
        await ws.send("Heartbeat")


async def player(player_id, uri):
    async with websockets.connect(uri, subprotocols=["echo"]) as ws:
        print(f"Player {player_id} connected to WebSocket server")

        try:
            # 启动心跳消息发送任务
            heartbeat_task = asyncio.create_task(send_heartbeat(ws))

            while True:
                # 异步接收消息
                message = await ws.recv()
                print(f"Player {player_id} received: {message}")

                # 异步处理消息
                await handle_message(player_id, ws, message)

        except websockets.exceptions.ConnectionClosedError:
            print(f"Player {player_id} disconnected")


async def main():
    uri = "ws://localhost:6666"
    players = 1  # 设置玩家数量

    # 创建多个玩家
    tasks = [player(i, uri) for i in range(players)]

    # 等待所有玩家完成
    await asyncio.gather(*tasks)


if __name__ == "__main__":
    asyncio.run(main())
