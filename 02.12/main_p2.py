def main():
    with open("input", "r") as file:
        result = 0
        for game in file:
            game_id, data = game.lstrip('Game ').split(":")

            draw_data = {"red": 0, "green": 0, "blue": 0}
            for draw in data.split(";"):
                for color in draw.split(","):
                    times, name = color.strip(' \n').split(' ')
                    if int(times) > draw_data[name]:
                        draw_data[name] = int(times)
            result += draw_data["red"] * draw_data["green"] * draw_data["blue"]

        print(result)


if __name__ == "__main__":
    main()
