import cv2


def remove_watermark(video_path, watermark_path):
    # 打开视频文件
    video = cv2.VideoCapture(video_path)

    # 读取水印图像
    watermark = cv2.imread(watermark_path)

    # 获取水印图像的宽度和高度
    w, h, _ = watermark.shape

    # 获取视频的帧率和尺寸
    fps = video.get(cv2.CAP_PROP_FPS)
    width = int(video.get(cv2.CAP_PROP_FRAME_WIDTH))
    height = int(video.get(cv2.CAP_PROP_FRAME_HEIGHT))

    # 创建输出视频文件
    output = cv2.VideoWriter("output.mp4", cv2.VideoWriter_fourcc(*"mp4v"), fps, (width, height))

    while True:
        # 读取视频的一帧
        ret, frame = video.read()

        if not ret:
            break

        # 在当前帧中定位水印的位置
        result = cv2.matchTemplate(frame, watermark, cv2.TM_CCOEFF_NORMED)
        _, _, _, max_loc = cv2.minMaxLoc(result)

        # 将水印从当前帧中去除
        frame[max_loc[1] : max_loc[1] + h, max_loc[0] : max_loc[0] + w] = [255, 255, 255]

        # 将处理后的帧写入输出视频文件
        output.write(frame)

    # 释放资源
    video.release()
    output.release()


# 调用函数进行水印去除
remove_watermark("a.mp4", "v.png")
