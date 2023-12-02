valid_game_data = {
    "red": 12,
    "green": 13,
    "blue": 14
}


def is_draw_valid(draw):
    return (
            draw["red"] <= valid_game_data["red"] and
            draw["green"] <= valid_game_data["green"] and
            draw["blue"] <= valid_game_data["blue"]
    )


def main():
    valid_game_ids = []
    with open("input", "r") as file:
        for game in file:
            is_game_valid = True
            game_id, data = game.lstrip('Game ').split(":")

            for draw in data.split(";"):
                draw_data = {"red": 0, "green": 0, "blue": 0}
                for color in draw.split(","):
                    times, name = color.strip(' \n').split(' ')
                    draw_data[name] += int(times)

                if not is_draw_valid(draw_data):
                    is_game_valid = False
                    break

            if is_game_valid:
                valid_game_ids.append(int(game_id))

        print(sum(valid_game_ids))


if __name__ == "__main__":
    main()
