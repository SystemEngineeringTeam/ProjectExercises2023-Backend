import sys


# controllerフォルダの中に入っているGenerateGraphPresenterをimportする。
# このとき、controllerフォルダの中に__init__.pyを作成しておく必要がある。
# __init__.pyを作らない場合は、controller.GenerateGraphPresenterと書く必要がある。
import GenerateGraphPresenter as GGP

def main():
    args = sys.argv

    try:
        boardId = int(args[1])
    except:
        print('boardIdを指定してください')
        return

    generateGraphPresenter = GGP.GenerateGraphPresenter(boardId)
    generateGraphPresenter.generateGraph()
    print("success")


if __name__ == '__main__':
    main()