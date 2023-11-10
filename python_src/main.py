import sys

from python_src.presenter.GenerateGraphPresenter import GenerateGraphPresenter


def main():
    args = sys.argv

    try:
        boardId = int(args[1])
    except:
        print('boardIdを指定してください')
        return

    generateGraphPresenter = GenerateGraphPresenter(boardId)
    generateGraphPresenter.generateGraph()
    print("success")


if __name__ == '__main__':
    main()