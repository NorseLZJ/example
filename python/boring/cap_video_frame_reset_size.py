import cv2
import numpy as np
import sys

if len(sys.argv) < 2:
    print(f"usage: {sys.argv[0]} videoFile")
    exit()

cap = cv2.VideoCapture(sys.argv[1])


# 检查视频是否成功打开
if not cap.isOpened():
    print("Error: Failed to open video file.")
    exit()

# 初始化帧计数器
frame_count = 0

# 目标分辨率
target_width = 1080
target_height = 720

# 每 15 帧截取一张图片
while cap.isOpened():
    ret, frame = cap.read()
    if not ret:
        break

    # 每 15 帧截取一张图片
    if frame_count % 15 == 0:
        # 获取原始图像大小
        original_height, original_width = frame.shape[:2]
        # 计算调整比例
        ratio = min(target_width / original_width, target_height / original_height)
        new_width = int(original_width * ratio)
        new_height = int(original_height * ratio)

        # 按比例缩放图像
        resized_frame = cv2.resize(frame, (new_width, new_height), interpolation=cv2.INTER_AREA)

        # 创建一个黑色背景的图像
        background = np.zeros((target_height, target_width, 3), dtype=np.uint8)

        # 计算图像放置位置
        x_offset = (target_width - new_width) // 2
        y_offset = (target_height - new_height) // 2

        # 将缩放后的图像叠加在黑色背景上
        background[y_offset:y_offset+new_height, x_offset:x_offset+new_width] = resized_frame

        # 指定保存图片的路径和文件名
        output_file = f"output_frame_{frame_count // 15}.jpg"

        # 保存调整后的图片
        cv2.imwrite(output_file, background)

    frame_count += 1

# 释放资源
cap.release()
