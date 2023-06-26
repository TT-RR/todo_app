import git
from datetime import datetime

# GitHubリポジトリのパスを指定します
repo_path = 'C:\Users\user\source\git-file\kusahayashi'

# コミットメッセージに現在の日時を追加するために、日時を取得します
current_datetime = datetime.now().strftime('%Y-%m-%d %H:%M:%S')

# リポジトリを開き、コミットとプッシュを実行します
repo = git.Repo(repo_path)
repo.git.add(all=True)
repo.git.commit('-m', f'自動プッシュ: {current_datetime}')
repo.git.push()
