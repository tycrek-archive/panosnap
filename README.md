# panosnap
Easily take Minecraft main menu panoramas

## Tutorial

1. Download the file linked below (eventually this will be on GitHub)
2. Set Minecraft to open at **922x922**
    - In the vanilla launcher this is done by going to **Installations** > **Edit** > **Resolution**
    - In MultiMC, go to **Edit Instance** > **Settings** > **Game Windows** and set Game Windows to `true`.
3. Run that installation and change the following in game options:
    - **FOV**: 90
    - Video Settings > **FOV Effects**: Off (this can be found in 1.16.2 and possibly other versions)
    - Video Settings > **Dynamic FOV**: Off (this is an Optifine setting)
4. Open the world or server you want a panorama from
5. If possible, disable `doDaylightCycle` gamerule (`/gamerule doDaylightCycle false`)
6. Run the program (should be fine to double click, running in command line will output any potential errors)
7. You will find the screenshots in your screenshots folder. Rename them from `panorama_0.png` to `panorama_5.png` (panorama_5 should be looking down)
8. Download the sample pack linked below and unzip it. Replace the files in `assets\minecraft\textures\gui\title\background\` with your screenshots. **Do not** delete `panorama_overlay.png`
9. Go back to the extracted folder and select `assets/` and `pack.mcmeta` (and optionally `pack.png` if you added one). Now add it to a zip by right clicking:
    - Windows' built-in right-click zip tool *should* be fine, but it is recommended to use 7-zip:
    - 7-zip > Add to archive. Set Format to `zip` and Compression Level to `store`. Name it and click Ok.
10. Place this new zip in `.minecraft/resourcepacks` and enjoy your panorama!

## Known issues
- Sometimes it can take the screenshots a bit *too* fast and chunks may not have loaded. I'll probably let you configure how fast it takes in the future to let chunks load. ~~Or just get a faster computer~~
- Shaders can have animations that might cause stitching issues. Right now I have no way to fix this.