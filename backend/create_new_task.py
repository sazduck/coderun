import os
import re


def create_next_task():
    # Регулярка для поиска папок типа task001
    pattern = re.compile(r"^task(\d{3})$")

    # Сканируем текущую директорию
    tasks = [0]  # Начинаем с 0 на случай, если папок еще нет
    for entry in os.listdir("."):
        if os.path.isdir(entry):
            match = pattern.match(entry)
            if match:
                tasks.append(int(match.group(1)))

    # Определяем номер следующей задачи
    next_num = max(tasks) + 1
    next_task_name = f"task{next_num:03d}"

    # Создаем папку
    os.makedirs(next_task_name, exist_ok=True)

    # Содержимое main.go
    main_content = """package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	Run(os.Stdin, os.Stdout)
}

func Run(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	next := func() int {
		s.Scan()
		val, _ := strconv.Atoi(s.Text())
		return val
	}

	val := next()

	b := bufio.NewWriterSize(w, 1<<20)
	defer b.Flush()
	b.WriteString(strconv.Itoa(val) + "\\n")
	return nil
}
"""

    # Содержимое main_test.go
    test_content = f"""package main_test

import (
	"backend/internal/testutil"
	{next_task_name} "backend/{next_task_name}"
	"testing"
)

func TestRun(t *testing.T) {{
	testutil.RunWithDefaultTestCasesPath(t, {next_task_name}.Run)
}}
"""

    # Записываем файлы
    files = {
        "main.go": main_content,
        "main_test.go": test_content,
        "test_cases.json": "[]",  # Пустой массив для корректной работы json.Unmarshal
    }

    for filename, content in files.items():
        with open(os.path.join(next_task_name, filename), "w", encoding="utf-8") as f:
            f.write(content)

    print(f"✅ Создана задача {next_task_name}:")
    print(f"   - {next_task_name}/main.go")
    print(f"   - {next_task_name}/main_test.go")
    print(f"   - {next_task_name}/test_cases.json")


if __name__ == "__main__":
    create_next_task()
