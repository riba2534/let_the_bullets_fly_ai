import json
import os


def extract_chinese_subtitles(filename):
    with open(filename, "r", encoding="utf-8") as file:
        content = file.read()

    subtitles = []
    subtitle_blocks = content.split("\n\n")
    for block in subtitle_blocks:
        lines = block.strip().split("\n")
        if len(lines) >= 3:
            subtitle_id = int(lines[0].strip())
            chinese_subtitle = lines[2].strip().split("\n")[0]
            subtitles.append({"id": subtitle_id, "content": chinese_subtitle})

    return subtitles


input_path = os.path.join('..', 'data', 'Let.the.Bullets.Fly.srt')
subtitles = extract_chinese_subtitles(input_path)

output_path = os.path.join('..', 'data', 'Let.the.Bullets.Fly.srt.json')
with open(output_path, "w", encoding="utf-8") as outfile:
    json.dump(subtitles, outfile, ensure_ascii=False, indent=4)

print("提取完成并保存到 Let.the.Bullets.Fly.srt.json 文件中")
