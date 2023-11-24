from faker import Faker


def generate_chinese_names(count):
    fake = Faker(locale="zh_CN")
    names = [fake.name() for _ in range(count)]
    return names


if __name__ == "__main__":
    chinese_names = generate_chinese_names(10)  # 生成10个虚构中文名字
    for name in chinese_names:
        print(f"刘{name[1:]}")
