import cv2
from PIL import Image
import numpy as np
import time


def video_frames_to_jpg(link: str):
    start = int(time.time())
    cap = cv2.VideoCapture(link)
    # 视频信息获取
    frame_count = cap.get(cv2.CAP_PROP_FRAME_COUNT)
    imageNum = 0
    sum = 0
    timef = int(frame_count / 9)  # 隔15帧保存一张图片
    # timef = 100  # 隔15帧保存一张图片
    if cap.isOpened:
        while True:
            sum += 1
            (frameState, frame) = cap.read()  # 记录每帧及获取状态
            if frameState == True and (sum % timef == 0):
                # 格式转变，BGRtoRGB
                frame = cv2.cvtColor(frame, cv2.COLOR_BGR2RGB)
                # 转变成Image
                frame = Image.fromarray(np.uint8(frame))
                frame = np.array(frame)
                # RGBtoBGR满足opencv显示格式
                frame = cv2.cvtColor(frame, cv2.COLOR_RGB2BGR)
                imageNum = imageNum + 1
                fileName = "image/image" + str(imageNum) + ".jpg"  # 存储路径
                cv2.imwrite(fileName, frame, [cv2.IMWRITE_JPEG_QUALITY, 100])
                print(fileName + " successfully write in")  # 输出存储状态
            elif frameState == False:
                break
    print("finish!")

    end = int(time.time())
    print("耗时:", end - start)
    cap.release()


if __name__ == "__main__":
    video_frames_to_jpg("C:\\Users\\Administrator\\Downloads\\8k.mkv")
    """
    proc_list = []
    for r, d, files in os.walk("image"):
        for file in files:
            if file.endswith("mp4") or file.endswith("mkv"):
                padre = os.path.join(r, file)
                padre = padre.replace("\\", "/")
                    proc_list.append((padre, r))
    for i, tp in tqdm(enumerate(proc_list)):
        print("Now processing video: " + tp[0])
        gen_captura(tp[0], tp[1])
    """
