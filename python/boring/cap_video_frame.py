import cv2
import sys

if len(sys.argv) < 2:
    print(f"usage: {sys.argv[0]} videoFile")
    exit()

cap = cv2.VideoCapture(sys.argv[1])

if not cap.isOpened():
    print("Error: Failed to open video file.")
    exit()

# 初始化帧计数器
frame_count = 0

# 每 15 帧截取一张图片
while cap.isOpened():
    ret, frame = cap.read()
    if not ret:
        break

    # 每 15 帧截取一张图片
    if frame_count % 60 == 0:
        output_file = f"img/output_frame_{frame_count // 60}.jpg"
        cv2.imwrite(output_file, frame)

    frame_count += 1

cap.release()
