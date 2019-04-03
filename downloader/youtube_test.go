package downloader

import (
	"testing"
)

func TestParseFileName(t *testing.T) {
	tests := []struct {
		ss    []string
		right string
	}{
		{[]string{
			"[youtube] lvObk195BQU: Downloading webpage",
			"[youtube] lvObk195BQU: Downloading video info webpage",
			"[youtube] lvObk195BQU: Extracting video information",
			"[youtube] lvObk195BQU: Downloading thumbnail ...",
			"[youtube] lvObk195BQU: Writing thumbnail to: Minions funny Memorable Moments movies and clips HD (episode 04)-lvObk195BQU.jpg",
			"[download] Destination: Minions funny Memorable Moments movies and clips HD (episode 04)-lvObk195BQU.mp4",
			"[download] 100% of 22.40MiB in 00:06",
		}, "Minions funny Memorable Moments movies and clips HD (episode 04)-lvObk195BQU.mp4"},
		{[]string{"[youtube] UqbTLJ0U84M: Downloading webpage",
			"[youtube] UqbTLJ0U84M: Downloading video info webpage",
			"[youtube] UqbTLJ0U84M: Extracting video information",
			"[info] Writing video description to: Men in Black (1997) - It's a Squid Scene (4_8) _ Movieclips-UqbTLJ0U84M.description",
			"[youtube] UqbTLJ0U84M: Downloading thumbnail ...",
			"[youtube] UqbTLJ0U84M: Writing thumbnail to: Men in Black (1997) - It's a Squid Scene (4_8) _ Movieclips-UqbTLJ0U84M.jpg",
			"[download] Men in Black (1997) - It's a Squid Scene (4_8) _ Movieclips-UqbTLJ0U84M.mp4 has already been downloaded",
			"[download] 100% of 18.79MiB"},
			"Men in Black (1997) - It's a Squid Scene (4_8) _ Movieclips-UqbTLJ0U84M.mp4"},
		{
			[]string{
				"[youtube] 4OSH3dYpnvo: Downloading webpage",
				"[youtube] 4OSH3dYpnvo: Downloading video info webpage",
				"[youtube] 4OSH3dYpnvo: Extracting video information",
				"[info] Writing video description to: 超人特攻隊2 _ HD最新中文電影預告 (The Incredibles 2)-4OSH3dYpnvo.description",
				"[youtube] 4OSH3dYpnvo: Downloading thumbnail ...",
				"[youtube] 4OSH3dYpnvo: Writing thumbnail to: 超人特攻隊2 _ HD最新中文電影預告 (The Incredibles 2)-4OSH3dYpnvo.jpg",
				"[download] Destination: 超人特攻隊2 _ HD最新中文電影預告 (The Incredibles 2)-4OSH3dYpnvo.mp4",
				"[download] 100% of 22.28MiB in 00:1220MiB/s ETA 00:000nown ETA",
			},
			"超人特攻隊2 _ HD最新中文電影預告 (The Incredibles 2)-4OSH3dYpnvo.mp4",
		},
	}
	for _, ts := range tests {
		title := ParseFileName(ts.ss)
		if title != ts.right {
			t.Errorf("title exptcted: [%s], but got [%s]", ts.right, title)
		}
	}

}

func TestDownload(t *testing.T) {
	DownloadCommand = "./bin/youtube.sh"
	ok, filename := Download("https://www.youtube.com/watch?v=UqbTLJ0U84M")
	if !ok || filename.FilePath == "" {
		t.Error("下载失败了")
	}

	//info, err := os.Stat(filename)

}
