import os
from striprtf.striprtf import rtf_to_text

def convert_rtf_to_txt(root_folder, output_folder):
    for dirpath, _, filenames in os.walk(root_folder):
        for filename in filenames:
            if filename.lower().endswith(".rtf"):
                input_path = os.path.join(dirpath, filename)

                # Чтение и конвертация RTF → Text
                with open(input_path, "r", encoding="utf-8", errors="ignore") as file:
                    rtf_content = file.read()
                    plain_text = rtf_to_text(rtf_content)

                # Вычисляем путь для сохранения .txt файла
                relative_path = os.path.relpath(dirpath, root_folder)
                output_dir = os.path.join(output_folder, relative_path)
                os.makedirs(output_dir, exist_ok=True)

                txt_filename = os.path.splitext(filename)[0] + ".txt"
                output_path = os.path.join(output_dir, txt_filename)

                # Сохраняем текст
                with open(output_path, "w", encoding="utf-8") as out_file:
                    out_file.write(plain_text)

                print(f"✅ {input_path} → {output_path}")

# 🔧 Пример использования
input_directory = "./data/ГОСТ Р ТК 164"
output_directory = "./data/converted_txt"

convert_rtf_to_txt(input_directory, output_directory)
